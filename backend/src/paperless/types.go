package paperless

type paperlessResponseResult interface {
	PaperlessDocument | PaperlessCorrespondent
}

type paperlessResponse interface {
	paperlessDocumentResponse | paperlessCorrespondentResponse
}

type PaperlessDocument struct {
	ID              int           `json:"id"`
	CorrespondentID int           `json:"correspondent"`
	Title           string        `json:"title"`
	Content         string        `json:"content"`
	Tags            []int         `json:"tags"`
	Created_date    PaperlessTime `json:"created_date"`
}

type PaperlessCorrespondent struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type paperlessDocumentResponse struct {
	Count    int                 `json:"count"`
	Next     string              `json:"next,omitempty"`
	Previous string              `json:"previous,omitempty"`
	All      []int               `json:"all"`
	Results  []PaperlessDocument `json:"results"`
}

type paperlessCorrespondentResponse struct {
	Count    int                      `json:"count"`
	Next     string                   `json:"next,omitempty"`
	Previous string                   `json:"previous,omitempty"`
	All      []int                    `json:"all"`
	Results  []PaperlessCorrespondent `json:"results"`
}
