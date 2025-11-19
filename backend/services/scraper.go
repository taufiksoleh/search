package services

import (
	"net/http"
	"strings"
	"sync"
	"time"

	"github.com/PuerkitoBio/goquery"
	"github.com/search-engine/backend/models"
)

type ScraperService struct {
	httpClient *http.Client
}

func NewScraperService() *ScraperService {
	return &ScraperService{
		httpClient: &http.Client{
			Timeout: 10 * time.Second,
		},
	}
}

// EnrichResults scrapes content from search result URLs to provide more context
func (s *ScraperService) EnrichResults(results []models.SearchResult) []models.SearchResult {
	if len(results) == 0 {
		return results
	}

	// Limit concurrent scraping
	maxConcurrent := 5
	if len(results) < maxConcurrent {
		maxConcurrent = len(results)
	}

	var wg sync.WaitGroup
	enriched := make([]models.SearchResult, len(results))
	semaphore := make(chan struct{}, maxConcurrent)

	for i, result := range results {
		wg.Add(1)
		go func(idx int, r models.SearchResult) {
			defer wg.Done()
			semaphore <- struct{}{}
			defer func() { <-semaphore }()

			content, err := s.scrapeContent(r.URL)
			if err == nil && content != "" {
				r.Content = content
			}
			enriched[idx] = r
		}(i, result)
	}

	wg.Wait()
	return enriched
}

// scrapeContent fetches and extracts main content from a URL
func (s *ScraperService) scrapeContent(url string) (string, error) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return "", err
	}

	// Set user agent to avoid blocks
	req.Header.Set("User-Agent", "Mozilla/5.0 (compatible; SearchBot/1.0)")

	resp, err := s.httpClient.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", nil
	}

	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		return "", err
	}

	// Remove script and style elements
	doc.Find("script, style, nav, footer, header, aside").Remove()

	// Try to find main content
	var content string

	// Try common content selectors
	selectors := []string{
		"article",
		"main",
		".content",
		".post-content",
		".article-content",
		"#content",
		".entry-content",
	}

	for _, selector := range selectors {
		selection := doc.Find(selector)
		if selection.Length() > 0 {
			content = s.extractText(selection)
			if len(content) > 200 {
				break
			}
		}
	}

	// Fallback to body if no specific content found
	if content == "" {
		content = s.extractText(doc.Find("body"))
	}

	// Limit content length
	if len(content) > 2000 {
		content = content[:2000] + "..."
	}

	return content, nil
}

// extractText cleans and extracts text from a selection
func (s *ScraperService) extractText(selection *goquery.Selection) string {
	text := selection.Text()

	// Clean up whitespace
	lines := strings.Split(text, "\n")
	var cleanLines []string

	for _, line := range lines {
		line = strings.TrimSpace(line)
		if line != "" {
			cleanLines = append(cleanLines, line)
		}
	}

	return strings.Join(cleanLines, " ")
}
