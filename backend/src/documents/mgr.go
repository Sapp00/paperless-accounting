package documents

import (
	"log"
	"sapp/paperless-accounting/config"
	"sapp/paperless-accounting/database"
	"sapp/paperless-accounting/paperless"

	"github.com/redis/go-redis/v9"
)

type DocumentMgr struct {
	paperless *paperless.Paperless
	db        *database.Queries
	client    *redis.Client
}

func NewManager(conf *config.Config) (*DocumentMgr, error) {
	p, err := paperless.Init(conf)

	if err != nil {
		return nil, err
	}

	client := redis.NewClient(&redis.Options{
		Addr:     conf.REDIS_ADDRESS,
		Password: conf.REDIS_PASSWORD, // no password set
		DB:       0,                   // use default DB
	})

	m := DocumentMgr{
		paperless: p,
		client:    client,
	}

	err = m.setupDB()
	if err != nil {
		return nil, err
	}
	log.Print("Set up DB successful")

	return &m, nil
}
