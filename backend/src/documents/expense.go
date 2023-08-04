package documents

import (
	"context"
	"sapp/paperless-accounting/paperless"
)

type Expense struct {
	// from database
	Date     paperless.Paperless
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

func (m *DocumentMgr) GetExpensesInYear(year string) ([]Expense, error) {
	p_result, err := m.paperless.GetDocuments(paperless.Expense, year)
	if err != nil {
		return nil, err
	}

	ctx := context.Background()
	db_result, err := m.db.ListExpenses(ctx)
	if err != nil {
		return nil, err
	}

	out := make([]Expense, len(p_result))
	l := 0
	i := 0
	for _, r := range p_result {
		for j := i; j < len(db_result); j++{
			if 
		}
	}

	return p_result, err
}
