package models

// SearchRequest represents the incoming search request
type SearchRequest struct {
	Query string `json:"query" binding:"required"`
}

// SearchResult represents a single search result
type SearchResult struct {
	Title       string `json:"title"`
	URL         string `json:"url"`
	Description string `json:"description"`
	Content     string `json:"content,omitempty"`
}

// SearchResponse represents the complete search response
type SearchResponse struct {
	Query      string         `json:"query"`
	Answer     string         `json:"answer"`
	Sources    []SearchResult `json:"sources"`
	RelatedQs  []string       `json:"related_questions,omitempty"`
}

// AIRequest represents a request to the AI service
type AIRequest struct {
	Query    string         `json:"query"`
	Context  []SearchResult `json:"context"`
}

// AIResponse represents the AI-generated response
type AIResponse struct {
	Answer    string   `json:"answer"`
	RelatedQs []string `json:"related_questions"`
}
