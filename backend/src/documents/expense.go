package documents

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"log"
	"sapp/paperless-accounting/database"
	"sapp/paperless-accounting/paperless"
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

func (m *DocumentMgr) GetExpensesInYear(year int) ([]Expense, error) {
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

	var out []Expense
	i := 0
paper_loop:
	for _, e_paper := range p_result {
		found := false
		fmt.Printf("Searching for %v\n", e_paper.ID)
		for ; i < len(db_result); i++ {
			e_db := db_result[i]

			if e_db.ID == int64(e_paper.ID) {

				// not the right year
				if year != 0 && e_db.Expensedate.Year() != year {
					fmt.Printf("doesnt match\n")
					continue paper_loop
				}

				// check validity
				e_price := float64(0)
				if e_db.Price.Valid {
					e_price = e_db.Price.Float64
				}

				// merge data
				expense := Expense{
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
			// create expense because none has been found
			_, err := m.db.CreateExpense(ctx, database.CreateExpenseParams{
				ID:          int64(e_paper.ID),
				Price:       sql.NullFloat64{},
				Expensedate: e_paper.Created_date.Time,
			})
			if err != nil {
				return nil, err
			}
			log.Printf("Added entry %d to database\n", e_paper.ID)
			/*
				expense := Expense{
					Date:          *paperless.NewPaperlessTime(time.Unix(0, 0)),
					Value:         float64(0),
					PaperlessID:   e_paper.ID,
					Correspondent: e_paper.CorrespondentID,
					Title:         e_paper.Title,
					Content:       e_paper.Content,
					Tags:          e_paper.Tags,
					Created_date:  e_paper.Created_date,
				}
				// add to output and set start for next loop
				out = append(out, expense)*/
		}
	}

	fmt.Printf("out: %v\n", out)

	return out, err
}

func (m *DocumentMgr) GetExpenses() ([]Expense, error) {
	return m.GetExpensesInYear(0)
}
