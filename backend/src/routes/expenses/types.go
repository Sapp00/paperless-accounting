package expenses

import (
	"sapp/paperless-accounting/config"
	"sapp/paperless-accounting/documents"
	"sapp/paperless-accounting/paperless"
)

type ExpenseRouter struct {
	conf *config.Config
	dm   *documents.DocumentMgr
}

type chartEntry struct {
	Date     paperless.PaperlessTime `json:"date"`
	Category string                  `json:"category"`
	Value    float32                 `json:"value"`
}
