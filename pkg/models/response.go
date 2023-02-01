package models

type Response struct {
	Data ResponseData `json:"data"`
}

type ResponseData struct {
	Results ResponseResults `json:"results"`
}

type ResponseResults struct {
	GoogleResults []Results `json:"google_results,omitempty"`
	RedditResults []Results `json:"reddit_results,omitempty"`
}

type Results struct {
	Link string `json:"link"`
}
