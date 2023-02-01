package models

type GoogleResponse struct {
	Kind              string
	Url               string
	Queries           string
	Context           string
	SearchInformation string
	Items             []GoogleItem
}

type GoogleItem struct {
	Link string
}
