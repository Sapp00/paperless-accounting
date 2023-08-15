package payments

import (
	"sapp/paperless-accounting/config"
	"sapp/paperless-accounting/documents"
)

type PaymentRouter struct {
	conf *config.Config
	dm   *documents.DocumentMgr
}

func New(conf *config.Config, dm *documents.DocumentMgr) (*PaymentRouter, error) {

	r := PaymentRouter{
		conf: conf,
		dm:   dm,
	}

	return &r, nil
}
