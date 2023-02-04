package controllers

import (
	"fmt"
	"os"
	"strconv"
	"universal-search/internal/models"
	"universal-search/internal/utils"
)

func GoogleSearch(query string, limit string) []models.Results {
	var googleKey = os.Getenv("GOOGLE_KEY")
	var cx = os.Getenv("CX")

	var limitStr = limit

	limitInt, err := strconv.Atoi(limit)

	if err == nil {
		if limitInt > 10 {
			limitStr = "10"
		}
	}

	var googleSearchURL = fmt.Sprintf("https://www.googleapis.com/customsearch/v1?key=" + googleKey + "&cx=" + cx + "&q=" + query + "&num=" + limitStr)

	var googleSearchResponse models.GoogleResponse
	utils.GetJson(googleSearchURL, &googleSearchResponse)
	var googleLinks = googleSearchResponse.Items

	var response []models.Results

	for i := 0; i < len(googleLinks); i++ {
		response = append(response, models.Results{Link: googleLinks[i].Link})
	}

	return response
}

func RedditSearch(query string, limit string) []models.Results {
	var limitStr = limit

	limitInt, err := strconv.Atoi(limit)

	if err == nil {
		if limitInt > 10 {
			limitStr = "10"
		}
	}

	var redditSearchURL = fmt.Sprintf("https://www.reddit.com/search.json?q=" + query + "&limit=" + limitStr)

	var redditSearchResponse models.RedditResponse
	utils.GetJson(redditSearchURL, &redditSearchResponse)
	var redditLinks = redditSearchResponse.Data.Children

	var response []models.Results

	for i := 0; i < len(redditLinks); i++ {
		response = append(response, models.Results{Link: redditLinks[i].Data.URL})
	}

	return response
}
