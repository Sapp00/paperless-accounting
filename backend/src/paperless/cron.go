package paperless

import (
	"context"
	"fmt"
	"log"
	"sapp/paperless-accounting/config"
	"sapp/paperless-accounting/crons"
	"time"

	"github.com/redis/go-redis/v9"
)

type PaperlessCron struct {
	client *redis.Client
}

func (p PaperlessCron) Run() {
	fmt.Printf("heyo")
}

func (p PaperlessCron) Interval() time.Duration {
	return time.Minute * 5
}

func StartCron(conf *config.Config) *crons.Cronjob {
	ctx := context.Background()

	client := redis.NewClient(&redis.Options{
		Addr:     conf.REDIS_ADDRESS,
		Password: conf.REDIS_PASSWORD, // no password set
		DB:       0,                   // use default DB
	})

	if pong := client.Ping(ctx); pong.String() != "ping: PONG" {
		log.Fatal(pong)
	}

	p := PaperlessCron{
		client: client,
	}
	c := crons.Start(p)

	err := client.Set(ctx, "foo", "bar", time.Minute*5).Err()
	if err != nil {
		panic(err)
	}

	val, err := client.Get(ctx, "foo").Result()
	if err != nil {
		panic(err)
	}
	fmt.Println("foo\n", val)

	err = client.Del(ctx, "foo").Err()
	if err != nil {
		panic(err)
	}

	return c
}
