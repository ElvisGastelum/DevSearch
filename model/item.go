package model

type SearchResults struct {
	Items []item `json:"items"`
}

type item struct {
	Link    string `json:"link"`
	Snippet string `json:"snippet"`
	Title   string `json:"title"`
}
