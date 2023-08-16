package correspondents

import (
	"sapp/paperless-accounting/config"
	"sapp/paperless-accounting/documents"
)

type CorrespondentRouter struct {
	conf *config.Config
	dm   *documents.DocumentMgr
}

func New(conf *config.Config, dm *documents.DocumentMgr) (*CorrespondentRouter, error) {

	r := CorrespondentRouter{
		conf: conf,
		dm:   dm,
	}

	return &r, nil
}
