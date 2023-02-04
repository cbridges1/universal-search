package handlers

import (
	"net/http"
	"sync"
	"universal-search/internal/controllers"
	"universal-search/internal/models"
	"universal-search/internal/utils"
)

func SearchHandler(res http.ResponseWriter, req *http.Request) {

	if req.Method != "GET" {
		// Add the response return message
		HandlerMessage := []byte(`{
		 "success": false,
		 "message": "Check your HTTP method: Invalid HTTP method executed",
		}`)

		utils.ReturnJsonResponse(res, http.StatusMethodNotAllowed, HandlerMessage)
		return
	}

	var query = req.URL.Query().Get("query")

	var limit = req.URL.Query().Get("limit")
	if len(limit) == 0 {
		limit = "10"
	}

	var googleResults []models.Results
	var redditResults []models.Results

	var wg sync.WaitGroup

	wg.Add(2)

	go func() {
		googleResults = controllers.GoogleSearch(query, limit)
		defer wg.Done()
	}()

	go func() {
		redditResults = controllers.RedditSearch(query, limit)
		defer wg.Done()
	}()

	wg.Wait()

	// googleResults = controllers.GoogleSearch(query, limit)
	// redditResults = controllers.RedditSearch(query, limit)

	var results = models.Response{
		Data: models.ResponseData{
			Results: models.ResponseResults{
				GoogleResults: googleResults,
				RedditResults: redditResults,
			},
		},
	}

	utils.ReturnJsonResponseJ(res, http.StatusOK, results)
}
