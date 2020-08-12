package model

type SearchResults struct {
	Items []Item `json:"items"`
}

// Item represent the result to the call from Google's API
type Item struct {
	Link    string `json:"link"`
	Snippet string `json:"snippet"`
	Title   string `json:"title"`
}
