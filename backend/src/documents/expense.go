package documents

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"log"
	"sapp/paperless-accounting/database"
	"sapp/paperless-accounting/paperless"
	"time"
)

type Expense struct {
	// from database
	Date  paperless.PaperlessTime
	Value float64

	// from paperless
	PaperlessID   int
	Correspondent int
	Title         string
	Content       string
	Tags          []int
	Created_date  paperless.PaperlessTime
}

func (m *DocumentMgr) GetExpense(id int) (*Expense, error) {
	e, err := m.paperless.GetDocument(paperless.Expense, id)
	if err != nil {
		return nil, err
	}

	ctx := context.Background()
	e_db, err := m.db.GetExpense(ctx, int64(id))
	if err != nil {
		return nil, err
	}

	fmt.Printf("res: %v\n", e_db)

	e_price := float64(0)
	if e_db.Price.Valid {
		e_price = e_db.Price.Float64
	}

	// merge data
	return &Expense{
		Date:          *paperless.NewPaperlessTime(e_db.Expensedate),
		Value:         e_price,
		PaperlessID:   e.ID,
		Correspondent: e.CorrespondentID,
		Title:         e.Title,
		Content:       e.Content,
		Tags:          e.Tags,
		Created_date:  e.Created_date,
	}, nil

	return nil, errors.New("documents:expense: could not find the respective id")
}

func (m *DocumentMgr) GetExpensesBetween(fromTimeStr string, toTimeStr string) ([]*Expense, error) {
	p_result, err := m.paperless.GetDocuments(paperless.Expense)
	if err != nil {
		return nil, err
	}

	ctx := context.Background()
	db_result, err := m.db.ListExpenses(ctx)
	if err != nil {
		return nil, err
	}

	fmt.Printf("db: %v\n", db_result)

	var fromTime time.Time
	if fromTimeStr == "-1" {
		fromTime = time.Unix(0, 0)
	} else {
		fromTime, err = time.Parse(`"2006-01-02"`, fromTimeStr)
		if err != nil {
			return nil, err
		}
	}
	var toTime time.Time
	if toTimeStr == "0" {
		toTime = time.Now()
	} else {
		toTime, err = time.Parse(`"2006-01-02"`, toTimeStr)
		if err != nil {
			return nil, err
		}
	}

	var out []*Expense
	i := 0
paper_loop:
	for _, e_paper := range p_result {
		found := false
		fmt.Printf("Searching for %v\n", e_paper.ID)
		for ; i < len(db_result); i++ {
			e_db := db_result[i]

			if e_db.ID == int64(e_paper.ID) {

				// not the right year
				if e_db.Expensedate.After(fromTime) && e_db.Expensedate.Before(toTime) {
					continue paper_loop
				}

				// check validity
				e_price := float64(0)
				if e_db.Price.Valid {
					e_price = e_db.Price.Float64
				}

				// merge data
				expense := &Expense{
					Date:          *paperless.NewPaperlessTime(e_db.Expensedate),
					Value:         e_price,
					PaperlessID:   e_paper.ID,
					Correspondent: e_paper.CorrespondentID,
					Title:         e_paper.Title,
					Content:       e_paper.Content,
					Tags:          e_paper.Tags,
					Created_date:  e_paper.Created_date,
				}
				// add to output and set start for next loop
				out = append(out, expense)
				i++
				found = true
				break
			}
			if e_db.ID > int64(e_paper.ID) {
				break
			}
		}
		// didnt find the element -> create a barebone
		if !found {
			// create expense because none has been found.
			// TODO: this should be moved to the cron!
			_, err := m.db.CreateExpense(ctx, database.CreateExpenseParams{
				ID:          int64(e_paper.ID),
				Price:       sql.NullFloat64{},
				Expensedate: e_paper.Created_date.Time,
			})
			if err != nil {
				return nil, err
			}
			log.Printf("Added entry %d to database\n", e_paper.ID)
		}
	}

	fmt.Printf("out: %v\n", out)

	return out, err
}

func (m *DocumentMgr) GetExpenses() ([]*Expense, error) {
	return m.GetExpensesBetween("-1", "0")
}

func (m *DocumentMgr) UpdateExpense(id int, date *paperless.PaperlessTime, value *float64) (*Expense, error) {
	exp, err := m.GetExpense(id)
	if err != nil {
		return nil, errors.New("expense cannot be updated because it does not exist. create it first")
	}
	if date != nil {
		exp.Date = *date
	}
	if value != nil {
		exp.Value = *value
	}

	ctx := context.Background()
	m.db.UpdateExpense(ctx, database.UpdateExpenseParams{Price: sql.NullFloat64{Valid: true, Float64: exp.Value}, Expensedate: exp.Date.Time, ID: int64(id)})

	return exp, nil
}
