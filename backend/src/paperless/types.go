package paperless

import "time"

type PaperlessDocument struct {
	Id            int           `json:"id"`
	Correspondent int           `json:"correspondent"`
	Title         string        `json:"title"`
	Content       string        `json:"content"`
	Tags          []int         `json:"tags"`
	Created_date  PaperlessTime `json:"created_date"`
}

type paperlessDocumentResponse struct {
	Count    int                 `json:"count"`
	Next     string              `json:"next,omitempty"`
	Previous string              `json:"previous"`
	All      []int               `json:"all"`
	Results  []PaperlessDocument `json:"results"`
}

type PaperlessTime struct {
	time.Time
}
