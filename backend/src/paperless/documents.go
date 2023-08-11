package paperless

import (
	"context"
	"encoding/json"
	"strconv"

	"github.com/redis/go-redis/v9"
)

type DocumentType uint8

const (
	Expense DocumentType = 0
	Income  DocumentType = 1
)

func (p *Paperless) GetDocuments(typ DocumentType) ([]PaperlessDocument, error) {
	ctx := context.Background()

	var tag string
	if typ == Expense {
		tag = "expenses"
	} else {
		tag = "incomes"
	}

	val, err := p.client.ZRange(ctx, tag, 0, -1).Result()

	if err != nil {
		if err.Error() == "redis: nil" {
			return nil, nil
		}
		return nil, err
	}

	var res []PaperlessDocument

	for _, v := range val {
		var p PaperlessDocument
		err = json.Unmarshal([]byte(v), &p)

		if err != nil {
			return nil, err
		}
		res = append(res, p)
	}

	return res, nil
}

func (p *Paperless) GetDocument(typ DocumentType, id int) (*PaperlessDocument, error) {
	ctx := context.Background()

	var tag string
	if typ == Expense {
		tag = "expenses"
	} else {
		tag = "incomes"
	}

	val, err := p.client.ZRangeByScore(ctx, tag, &redis.ZRangeBy{
		Min:   strconv.Itoa(id),
		Max:   strconv.Itoa(id),
		Count: 1,
	}).Result()

	if err != nil {
		if err.Error() == "redis: nil" {
			return nil, nil
		}
		return nil, err
	}

	var res PaperlessDocument

	err = json.Unmarshal([]byte(val[0]), &res)

	if err != nil {
		return nil, err
	}

	return &res, nil
}
