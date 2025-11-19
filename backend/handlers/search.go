package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/search-engine/backend/models"
	"github.com/search-engine/backend/services"
)

type SearchHandler struct {
	searchService  *services.SearchService
	aiService      *services.AIService
	scraperService *services.ScraperService
}

func NewSearchHandler(search *services.SearchService, ai *services.AIService, scraper *services.ScraperService) *SearchHandler {
	return &SearchHandler{
		searchService:  search,
		aiService:      ai,
		scraperService: scraper,
	}
}

func (h *SearchHandler) Search(c *gin.Context) {
	var req models.SearchRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request: query is required"})
		return
	}

	// Step 1: Get search results from search API
	searchResults, err := h.searchService.Search(req.Query)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch search results"})
		return
	}

	// Step 2: Scrape content from top results
	enrichedResults := h.scraperService.EnrichResults(searchResults)

	// Step 3: Generate AI answer based on search results
	aiResponse, err := h.aiService.GenerateAnswer(req.Query, enrichedResults)
	if err != nil {
		// Return results without AI answer if AI fails
		c.JSON(http.StatusOK, models.SearchResponse{
			Query:   req.Query,
			Answer:  "Unable to generate AI summary at this time.",
			Sources: enrichedResults,
		})
		return
	}

	// Step 4: Return complete response
	response := models.SearchResponse{
		Query:     req.Query,
		Answer:    aiResponse.Answer,
		Sources:   enrichedResults,
		RelatedQs: aiResponse.RelatedQs,
	}

	c.JSON(http.StatusOK, response)
}
