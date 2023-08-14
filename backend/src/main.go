package main

import (
	"fmt"
	"os"
	"sapp/paperless-accounting/config"
	"sapp/paperless-accounting/crons/fetchcron"
	"sapp/paperless-accounting/documents"
	"sapp/paperless-accounting/graph/gql_generated"
	"sapp/paperless-accounting/graph/gql_resolvers"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/gin-gonic/gin"
)

const defaultPort = "8080"

func graphqlHandler(doc *documents.DocumentMgr) gin.HandlerFunc {
	rslv := gql_resolvers.Resolver{Dm: doc}
	h := handler.NewDefaultServer(gql_generated.NewExecutableSchema(gql_generated.Config{Resolvers: &rslv}))
	return func(c *gin.Context) {
		h.ServeHTTP(c.Writer, c.Request)
	}
}

// Defining the Playground handler
func playgroundHandler() gin.HandlerFunc {
	h := playground.Handler("GraphQL", "/query")

	return func(c *gin.Context) {
		h.ServeHTTP(c.Writer, c.Request)
	}
}

func main() {
	conf, err := config.New()

	if err != nil {
		fmt.Printf("cannot load config: %s\n", err)
		os.Exit(1)
	}

	fetchcron.StartCron(conf)

	doc, err := documents.NewManager(conf)
	if err != nil {
		panic(err)
	}

	// Setting up Gin
	r := gin.Default()
	r.POST("/query", graphqlHandler(doc))
	r.GET("/", playgroundHandler())
	r.Run()
}
