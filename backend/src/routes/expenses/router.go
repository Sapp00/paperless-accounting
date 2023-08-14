package expenses

import (
	"sapp/paperless-accounting/config"
	"sapp/paperless-accounting/documents"
)

type ExpenseRouter struct {
	conf *config.Config
	dm   *documents.DocumentMgr
}

func New(conf *config.Config, dm *documents.DocumentMgr) (*ExpenseRouter, error) {

	r := ExpenseRouter{
		conf: conf,
		dm:   dm,
	}

	return &r, nil
}
