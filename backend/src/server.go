package main

import (
	"log"
	"net/http"
	"os"
	"sapp/paperless-accounting/config"
	"sapp/paperless-accounting/documents"
	"sapp/paperless-accounting/graph/gql_generated"
	"sapp/paperless-accounting/graph/gql_resolvers"
	"sapp/paperless-accounting/paperless"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
)

const defaultPort = "8080"

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	conf, err := config.New()
	if err != nil {
		panic(err)
	}

	paperless.StartCron(conf)

	doc, err := documents.NewManager(conf)
	if err != nil {
		panic(err)
	}

	rslv := gql_resolvers.Resolver{Dm: doc}
	srv := handler.NewDefaultServer(gql_generated.NewExecutableSchema(gql_generated.Config{Resolvers: &rslv}))

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
