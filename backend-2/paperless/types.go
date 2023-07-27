package paperless

type PaperlessDocument struct {
	Id            int      `json:"id"`
	Correspondent int      `json:"correspondent"`
	Title         string   `json:"title"`
	Content       string   `json:"content"`
	Tags          []string `json:"tags"`
	Created_date  string   `json:"created_date"`
}

type paperlessDocumentResponse struct {
	Count    int                 `json:"count"`
	Next     string              `json:"next"`
	Previous string              `json:"previous"`
	All      []int               `json:"all"`
	Results  []PaperlessDocument `json:"results"`
}
