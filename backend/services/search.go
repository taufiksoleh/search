package services

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"

	"github.com/search-engine/backend/models"
)

type SearchService struct {
	serpAPIKey      string
	braveAPIKey     string
	httpClient      *http.Client
}

func NewSearchService() *SearchService {
	return &SearchService{
		serpAPIKey:  os.Getenv("SERPAPI_KEY"),
		braveAPIKey: os.Getenv("BRAVE_SEARCH_API_KEY"),
		httpClient:  &http.Client{},
	}
}

// Search performs a web search and returns results
func (s *SearchService) Search(query string) ([]models.SearchResult, error) {
	// Try Brave Search first, then fall back to SerpAPI
	if s.braveAPIKey != "" {
		return s.searchBrave(query)
	}
	if s.serpAPIKey != "" {
		return s.searchSerp(query)
	}

	// Return mock results for demo if no API keys
	return s.mockSearch(query), nil
}

// searchBrave uses Brave Search API
func (s *SearchService) searchBrave(query string) ([]models.SearchResult, error) {
	apiURL := fmt.Sprintf("https://api.search.brave.com/res/v1/web/search?q=%s&count=10", url.QueryEscape(query))

	req, err := http.NewRequest("GET", apiURL, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Accept", "application/json")
	req.Header.Set("X-Subscription-Token", s.braveAPIKey)

	resp, err := s.httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var braveResp struct {
		Web struct {
			Results []struct {
				Title       string `json:"title"`
				URL         string `json:"url"`
				Description string `json:"description"`
			} `json:"results"`
		} `json:"web"`
	}

	if err := json.Unmarshal(body, &braveResp); err != nil {
		return nil, err
	}

	var results []models.SearchResult
	for _, r := range braveResp.Web.Results {
		results = append(results, models.SearchResult{
			Title:       r.Title,
			URL:         r.URL,
			Description: r.Description,
		})
	}

	return results, nil
}

// searchSerp uses SerpAPI for Google results
func (s *SearchService) searchSerp(query string) ([]models.SearchResult, error) {
	apiURL := fmt.Sprintf("https://serpapi.com/search.json?q=%s&api_key=%s&num=10",
		url.QueryEscape(query), s.serpAPIKey)

	resp, err := s.httpClient.Get(apiURL)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var serpResp struct {
		OrganicResults []struct {
			Title   string `json:"title"`
			Link    string `json:"link"`
			Snippet string `json:"snippet"`
		} `json:"organic_results"`
	}

	if err := json.Unmarshal(body, &serpResp); err != nil {
		return nil, err
	}

	var results []models.SearchResult
	for _, r := range serpResp.OrganicResults {
		results = append(results, models.SearchResult{
			Title:       r.Title,
			URL:         r.Link,
			Description: r.Snippet,
		})
	}

	return results, nil
}

// mockSearch returns mock results for demo purposes
func (s *SearchService) mockSearch(query string) []models.SearchResult {
	return []models.SearchResult{
		{
			Title:       fmt.Sprintf("Understanding %s - Complete Guide", query),
			URL:         "https://example.com/guide",
			Description: fmt.Sprintf("A comprehensive guide to understanding %s with examples and best practices.", query),
		},
		{
			Title:       fmt.Sprintf("%s - Wikipedia", query),
			URL:         "https://en.wikipedia.org/wiki/Example",
			Description: fmt.Sprintf("Wikipedia article about %s covering history, applications, and more.", query),
		},
		{
			Title:       fmt.Sprintf("How to Use %s Effectively", query),
			URL:         "https://example.com/tutorial",
			Description: fmt.Sprintf("Step-by-step tutorial on using %s in real-world scenarios.", query),
		},
		{
			Title:       fmt.Sprintf("%s Best Practices 2024", query),
			URL:         "https://example.com/best-practices",
			Description: fmt.Sprintf("Latest best practices and recommendations for %s.", query),
		},
		{
			Title:       fmt.Sprintf("Getting Started with %s", query),
			URL:         "https://example.com/getting-started",
			Description: fmt.Sprintf("Beginner-friendly introduction to %s.", query),
		},
	}
}
