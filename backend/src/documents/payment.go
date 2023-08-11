package documents

import "sapp/paperless-accounting/paperless"

type Payment struct {
	// from database
	ID        int
	Date      paperless.PaperlessTime
	Value     float32
	ExpenseID int
}
