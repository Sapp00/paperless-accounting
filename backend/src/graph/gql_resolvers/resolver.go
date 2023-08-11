package gql_resolvers

import "sapp/paperless-accounting/documents"

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	Dm *documents.DocumentMgr
}
