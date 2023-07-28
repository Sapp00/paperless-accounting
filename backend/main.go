package main

import (
	"fmt"
	"os"
	"sapp/paperless-accounting/config"
	"sapp/paperless-accounting/routes"
)

func main() {
	conf, err := config.New()

	if err != nil {
		fmt.Printf("cannot load config: %s\n", err)
		os.Exit(1)
	}

	routes, err := routes.New(conf)

	if err != nil {
		fmt.Printf("cannot setup server %s\n", err)
		os.Exit(1)
	}

	err = routes.Setup()

	if err != nil {
		fmt.Printf("cannot start server %s\n", err)
		os.Exit(1)
	}

}
