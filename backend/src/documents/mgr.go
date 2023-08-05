package documents

import (
	"log"
	"sapp/paperless-accounting/config"
	"sapp/paperless-accounting/database"
	"sapp/paperless-accounting/paperless"
)

type DocumentMgr struct {
	paperless *paperless.Paperless
	db        *database.Queries
}

func NewManager(conf *config.Config) (*DocumentMgr, error) {
	p, err := paperless.Init(conf)

	if err != nil {
		return nil, err
	}

	m := DocumentMgr{
		paperless: p,
	}

	err = m.setupDB()
	if err != nil {
		return nil, err
	}
	log.Print("Setted up DB")

	return &m, nil
}
