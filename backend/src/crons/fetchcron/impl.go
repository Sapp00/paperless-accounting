package fetchcron

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"sapp/paperless-accounting/config"
	"sapp/paperless-accounting/crons"
	"sapp/paperless-accounting/database"
	"sapp/paperless-accounting/documents"
	"sapp/paperless-accounting/paperless"
	"time"

	"github.com/redis/go-redis/v9"
)

type PaperlessCron struct {
	client    *redis.Client
	paperless *paperless.Paperless
	dm        *documents.DocumentMgr
	conf      *config.Config
	db        *database.Queries
}

func (p PaperlessCron) Run() {
	all_expenses, err := p.paperless.PaperlessDocumentQuery("tag:" + p.conf.PAPERLESS_EXPENSE_TAG)

	if err != nil {
		panic(err)
	}
	all_incomes, err := p.paperless.PaperlessDocumentQuery("tag:" + p.conf.PAPERLESS_INCOME_TAG)
	if err != nil {
		panic(err)
	}

	ctx := context.Background()
	for _, e := range all_expenses {

		e_db, err := p.db.GetExpense(ctx, int64(e.ID))
		if err != nil {
			// not found, create DB entry
			if err.Error() == "redis: nil" {
				e_db, err = p.db.CreateExpense(ctx, database.CreateExpenseParams{
					ID:          int64(e.ID),
					Price:       sql.NullFloat64{},
					Expensedate: e.Created_date.Time,
				})
				if err != nil {
					panic(err)
				}
			} else {
				panic(err)
			}
		}

		// check price
		e_price := float64(0)
		if e_db.Price.Valid {
			e_price = e_db.Price.Float64
		}

		exp := documents.Expense{
			Date:          *paperless.NewPaperlessTime(e_db.Expensedate),
			Value:         e_price,
			PaperlessID:   e.ID,
			Correspondent: e.CorrespondentID,
			Title:         e.Title,
			Content:       e.Content,
			Tags:          e.Tags,
			Created_date:  e.Created_date,
		}

		ej, err := json.Marshal(&exp)
		if err != nil {
			panic(err)
		}

		// store id
		err = p.client.ZAdd(ctx, "expenses", redis.Z{Score: float64(e.ID), Member: ej}).Err()
		if err != nil {
			panic(err)
		}
	}

	for _, e := range all_incomes {

		e_db, err := p.db.GetIncome(ctx, int64(e.ID))
		if err != nil {
			// not found, create DB entry
			if err.Error() == "redis: nil" {
				e_db, err = p.db.CreateIncome(ctx, database.CreateIncomeParams{
					ID:         int64(e.ID),
					Price:      sql.NullFloat64{},
					Incomedate: e.Created_date.Time,
				})
				if err != nil {
					panic(err)
				}
			} else {
				panic(err)
			}
		}

		// check price
		e_price := float64(0)
		if e_db.Price.Valid {
			e_price = e_db.Price.Float64
		}

		exp := documents.Expense{
			Date:          *paperless.NewPaperlessTime(e_db.Incomedate),
			Value:         e_price,
			PaperlessID:   e.ID,
			Correspondent: e.CorrespondentID,
			Title:         e.Title,
			Content:       e.Content,
			Tags:          e.Tags,
			Created_date:  e.Created_date,
		}

		ej, err := json.Marshal(&exp)
		if err != nil {
			panic(err)
		}

		// store id
		err = p.client.ZAdd(ctx, "incomes", redis.Z{Score: float64(e.ID), Member: ej}).Err()
		if err != nil {
			panic(err)
		}
	}

	// expiry time = 5min
	p.client.Expire(ctx, "expenses", time.Minute*5)
	p.client.Expire(ctx, "incomes", time.Minute*5)

	// TODO: previous code should only be executed initially. in the future, the written arrays should be replaced with JSON
	// the JSON then will then be written to the HSet and

	fmt.Printf("Have written %v expense & %v income records to redis.\n", len(all_expenses), len(all_incomes))
}

func (p PaperlessCron) Interval() time.Duration {
	return time.Minute * 5
}

func StartCron(conf *config.Config) (*crons.Cronjob, error) {
	ctx := context.Background()

	client := redis.NewClient(&redis.Options{
		Addr:     conf.REDIS_ADDRESS,
		Password: conf.REDIS_PASSWORD, // no password set
		DB:       0,                   // use default DB
	})

	if pong := client.Ping(ctx); pong.String() != "ping: PONG" {
		log.Fatal(pong)
	}

	pl, err := paperless.Init(conf)
	if err != nil {
		return nil, err
	}

	p := PaperlessCron{
		client:    client,
		paperless: pl,
		conf:      conf,
	}
	c := crons.Start(p)

	return c, nil
}
