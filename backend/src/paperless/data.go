package paperless

import (
	"context"
	"encoding/json"
)

type DocumentType uint8

const (
	Expense DocumentType = 0
	Income  DocumentType = 1
)

func (p *Paperless) GetDocuments(typ DocumentType, year string) ([]PaperlessDocument, error) {
	ctx := context.Background()

	var tag string
	if typ == Expense {
		tag = "expenses"
	} else {
		tag = "incomes"
	}

	val, err := p.client.HGet(ctx, tag, year).Result()

	if err != nil {
		return nil, err
	}

	var res []PaperlessDocument

	err = json.Unmarshal([]byte(val), &res)

	if err != nil {
		return nil, err
	}

	return res, nil
}
