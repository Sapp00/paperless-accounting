package paperless

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"sapp/paperless-accounting/config"
	"sapp/paperless-accounting/crons"
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

	ctx := context.Background()
	for _, e := range all_expenses {
		ej, err := json.Marshal(&e)
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
		ej, err := json.Marshal(&e)
		if err != nil {
			panic(err)
		}
		err = p.client.ZAdd(ctx, "incomes", redis.Z{Score: float64(e.ID), Member: ej}).Err()
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
