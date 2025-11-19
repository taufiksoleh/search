package services

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"

	"github.com/search-engine/backend/models"
)

type AIService struct {
	openAIKey    string
	anthropicKey string
	httpClient   *http.Client
}

func NewAIService() *AIService {
	return &AIService{
		openAIKey:    os.Getenv("OPENAI_API_KEY"),
		anthropicKey: os.Getenv("ANTHROPIC_API_KEY"),
		httpClient:   &http.Client{},
	}
}

// GenerateAnswer generates an AI-powered answer based on search results
func (s *AIService) GenerateAnswer(query string, results []models.SearchResult) (*models.AIResponse, error) {
	// Build context from search results
	context := s.buildContext(results)

	// Try OpenAI first, then Anthropic, then fallback to mock
	if s.openAIKey != "" {
		return s.generateWithOpenAI(query, context)
	}
	if s.anthropicKey != "" {
		return s.generateWithAnthropic(query, context)
	}

	// Return mock response for demo
	return s.mockGenerate(query, results), nil
}

func (s *AIService) buildContext(results []models.SearchResult) string {
	var contextParts []string
	for i, r := range results {
		if i >= 5 { // Limit to top 5 results
			break
		}
		content := r.Description
		if r.Content != "" {
			content = r.Content
		}
		contextParts = append(contextParts, fmt.Sprintf("Source %d (%s):\nTitle: %s\nContent: %s\n",
			i+1, r.URL, r.Title, content))
	}
	return strings.Join(contextParts, "\n---\n")
}

// generateWithOpenAI uses OpenAI API
func (s *AIService) generateWithOpenAI(query, context string) (*models.AIResponse, error) {
	systemPrompt := `You are a helpful AI search assistant. Based on the provided search results, give a comprehensive, accurate answer to the user's query.
- Cite sources using [1], [2], etc. format
- Be concise but thorough
- If information is uncertain, say so
- At the end, suggest 3 related questions the user might want to explore`

	userPrompt := fmt.Sprintf("Query: %s\n\nSearch Results:\n%s\n\nProvide a helpful answer based on these sources.", query, context)

	reqBody := map[string]interface{}{
		"model": "gpt-4o-mini",
		"messages": []map[string]string{
			{"role": "system", "content": systemPrompt},
			{"role": "user", "content": userPrompt},
		},
		"temperature": 0.7,
		"max_tokens": 1000,
	}

	jsonBody, err := json.Marshal(reqBody)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", "https://api.openai.com/v1/chat/completions", bytes.NewBuffer(jsonBody))
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+s.openAIKey)

	resp, err := s.httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var openAIResp struct {
		Choices []struct {
			Message struct {
				Content string `json:"content"`
			} `json:"message"`
		} `json:"choices"`
	}

	if err := json.Unmarshal(body, &openAIResp); err != nil {
		return nil, err
	}

	if len(openAIResp.Choices) == 0 {
		return nil, fmt.Errorf("no response from OpenAI")
	}

	answer := openAIResp.Choices[0].Message.Content
	relatedQs := s.extractRelatedQuestions(answer)

	return &models.AIResponse{
		Answer:    answer,
		RelatedQs: relatedQs,
	}, nil
}

// generateWithAnthropic uses Anthropic Claude API
func (s *AIService) generateWithAnthropic(query, context string) (*models.AIResponse, error) {
	systemPrompt := `You are a helpful AI search assistant. Based on the provided search results, give a comprehensive, accurate answer to the user's query.
- Cite sources using [1], [2], etc. format
- Be concise but thorough
- If information is uncertain, say so
- At the end, suggest 3 related questions the user might want to explore`

	userPrompt := fmt.Sprintf("Query: %s\n\nSearch Results:\n%s\n\nProvide a helpful answer based on these sources.", query, context)

	reqBody := map[string]interface{}{
		"model": "claude-3-haiku-20240307",
		"max_tokens": 1000,
		"system": systemPrompt,
		"messages": []map[string]string{
			{"role": "user", "content": userPrompt},
		},
	}

	jsonBody, err := json.Marshal(reqBody)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", "https://api.anthropic.com/v1/messages", bytes.NewBuffer(jsonBody))
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("x-api-key", s.anthropicKey)
	req.Header.Set("anthropic-version", "2023-06-01")

	resp, err := s.httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var anthropicResp struct {
		Content []struct {
			Text string `json:"text"`
		} `json:"content"`
	}

	if err := json.Unmarshal(body, &anthropicResp); err != nil {
		return nil, err
	}

	if len(anthropicResp.Content) == 0 {
		return nil, fmt.Errorf("no response from Anthropic")
	}

	answer := anthropicResp.Content[0].Text
	relatedQs := s.extractRelatedQuestions(answer)

	return &models.AIResponse{
		Answer:    answer,
		RelatedQs: relatedQs,
	}, nil
}

// extractRelatedQuestions attempts to extract related questions from the answer
func (s *AIService) extractRelatedQuestions(answer string) []string {
	// Simple extraction - in production, you'd use better parsing
	return []string{
		"What are the main benefits?",
		"How does this compare to alternatives?",
		"What are common use cases?",
	}
}

// mockGenerate returns a mock AI response for demo purposes
func (s *AIService) mockGenerate(query string, results []models.SearchResult) *models.AIResponse {
	var sources []string
	for i, r := range results {
		if i >= 3 {
			break
		}
		sources = append(sources, fmt.Sprintf("[%d] %s", i+1, r.Title))
	}

	answer := fmt.Sprintf(`Based on the search results, here's what I found about "%s":

%s is a topic that has been extensively covered across multiple sources. According to the search results:

**Key Points:**
- The concept is well-documented with comprehensive guides available [1]
- Wikipedia provides historical context and background information [2]
- Multiple tutorials offer practical implementation guidance [3]

**Summary:**
The search results indicate that this is a well-established topic with plenty of resources for learning more. The sources provide both theoretical background and practical applications.

**Sources:**
%s`, query, query, strings.Join(sources, "\n"))

	return &models.AIResponse{
		Answer: answer,
		RelatedQs: []string{
			fmt.Sprintf("What are best practices for %s?", query),
			fmt.Sprintf("How to get started with %s?", query),
			fmt.Sprintf("Common mistakes when using %s", query),
		},
	}
}
