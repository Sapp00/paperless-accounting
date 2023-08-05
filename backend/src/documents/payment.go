package documents

import "sapp/paperless-accounting/paperless"

type Payment struct {
	// from database
	Date      paperless.PaperlessTime
	Category  paperless.DocumentType
	Value     float32
	ExpenseID int
}
