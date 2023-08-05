package documents

import (
	"context"
	"sapp/paperless-accounting/paperless"
	"time"
)

type Income struct {
	// from database
	Date  *paperless.PaperlessTime
	Value float64

	// from paperless
	PaperlessID   int
	Correspondent int
	Title         string
	Content       string
	Tags          []int
	Created_date  paperless.PaperlessTime
}

func (m *DocumentMgr) GetIncomesInYear(year string) ([]Income, error) {
	p_result, err := m.paperless.GetDocuments(paperless.Income, year)
	if err != nil {
		return nil, err
	}

	ctx := context.Background()
	db_result, err := m.db.Listincomes(ctx)
	if err != nil {
		return nil, err
	}

	var out []Income
	i := 0
	for _, e_paper := range p_result {
		for ; i < len(db_result); i++ {
			e_db := db_result[i]
			if e_db.ID == int64(e_paper.ID) {
				// check validity
				e_time := time.Unix(0, 0)
				if e_db.Incomedate.Valid {
					e_time = e_db.Incomedate.Time
				}
				e_price := float64(0)
				if e_db.Price.Valid {
					e_price = e_db.Price.Float64
				}

				income := Income{
					Date:          paperless.NewPaperlessTime(e_time),
					Value:         e_price,
					PaperlessID:   e_paper.ID,
					Correspondent: e_paper.CorrespondentID,
					Title:         e_paper.Title,
					Content:       e_paper.Content,
					Tags:          e_paper.Tags,
					Created_date:  e_paper.Created_date,
				}
				out = append(out, income)
				break
			}
		}
	}

	return out, err
}
