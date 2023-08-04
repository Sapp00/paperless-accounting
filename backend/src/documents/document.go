package documents

import "sapp/paperless-accounting/paperless"

type Income struct {
	// from database
	Date     paperless.PaperlessTime
	Category paperless.DocumentType
	Value    float32

	// from paperless
	PaperlessID   int
	Correspondent int
	Title         string
	Content       string
	Tags          []int
	Created_date  paperless.PaperlessTime
}

type Payment struct {
	// from database
	Date      paperless.PaperlessTime
	Category  paperless.DocumentType
	Value     float32
	ExpenseID int
}
