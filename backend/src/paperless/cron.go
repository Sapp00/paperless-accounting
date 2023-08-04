package paperless

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"sapp/paperless-accounting/config"
	"sapp/paperless-accounting/crons"
	"strconv"
	"time"

	"github.com/redis/go-redis/v9"
)

type PaperlessCron struct {
	client    *redis.Client
	paperless *Paperless
	conf      *config.Config
}

func (p PaperlessCron) Run() {
	all_expenses, err := p.paperless.paperlessDocumentQuery("tag:" + p.conf.PAPERLESS_EXPENSE_TAG)

	if err != nil {
		panic(err)
	}
	all_incomes, err := p.paperless.paperlessDocumentQuery("tag:" + p.conf.PAPERLESS_INCOME_TAG)
	if err != nil {
		panic(err)
	}

	// group by year
	expense_years := map[int][]PaperlessDocument{}
	for _, e := range all_expenses {
		year := e.Created_date.Year()

		if arr, ok := expense_years[year]; ok {
			expense_years[year] = append(arr, e)
		} else {
			expense_years[year] = []PaperlessDocument{e}
		}
	}

	// group by year
	income_years := map[int][]PaperlessDocument{}
	for _, e := range all_incomes {
		year := e.Created_date.Year()

		if arr, ok := income_years[year]; ok {
			income_years[year] = append(arr, e)
		} else {
			income_years[year] = []PaperlessDocument{e}
		}
	}

	ctx := context.Background()

	// write to redis
	for k, v := range expense_years {
		field := strconv.Itoa(k)

		value, err := json.Marshal(v)
		if err != nil {
			panic(err)
		}

		err = p.client.HSet(ctx, "expenses", field, value).Err()
		if err != nil {
			panic(err)
		}
	}
	for k, v := range income_years {
		field := strconv.Itoa(k)

		value, err := json.Marshal(v)
		if err != nil {
			panic(err)
		}

		err = p.client.HSet(ctx, "incomes", field, value).Err()
		if err != nil {
			panic(err)
		}
	}

	// expiry time = 5min
	p.client.Expire(ctx, "expenses", time.Minute*5)
	p.client.Expire(ctx, "incomes", time.Minute*5)

	// Todo: previous code should only be executed initially. in the future, the written arrays should be replaced with JSON
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

	pl, err := Init(conf)
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
