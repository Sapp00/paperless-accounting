package paperless

import (
	"errors"
	"sapp/paperless-accounting/config"

	"github.com/redis/go-redis/v9"
)

type Paperless struct {
	conf           *config.Config
	client         *redis.Client
	correspondents map[int]*PaperlessCorrespondent
}

func Init(conf *config.Config) (*Paperless, error) {
	client := redis.NewClient(&redis.Options{
		Addr:     conf.REDIS_ADDRESS,
		Password: conf.REDIS_PASSWORD, // no password set
		DB:       0,                   // use default DB
	})
	p := Paperless{
		conf:   conf,
		client: client,
	}

	return &p, nil
}

func (p *Paperless) GetCorrespondent(id int) (*PaperlessCorrespondent, error) {
	if val, ok := p.correspondents[id]; ok {
		return val, nil
	}
	return nil, errors.New("cannot find the correspondent")
}
