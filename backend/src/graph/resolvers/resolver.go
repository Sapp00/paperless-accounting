package resolvers

import (
	"sapp/paperless-accounting/documents"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	dm *documents.DocumentMgr
}
