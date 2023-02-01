package models

type RedditResponse struct {
	Kind string
	Data RedditData
}

type RedditData struct {
	Children []RedditItem
}

type RedditItem struct {
	Data RedditItemData
}

type RedditItemData struct {
	URL string
}
