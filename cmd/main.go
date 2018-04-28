package main

import (
	"flag"
	"log"
	"github.com/Xzya/headless-cms/api"
	"github.com/Xzya/headless-cms"
	"github.com/Xzya/headless-cms/config"
	"fmt"
)

var (
	dialect      = flag.String("dialect", "postgres", "The database dialect")
	databaseHost = flag.String("db-host", "localhost", "The database hostname")
	databasePort = flag.String("db-port", "5432", "The database port")
	database     = flag.String("database", "cms", "The database name")
	user         = flag.String("user", "test", "The database username")
	pass         = flag.String("pass", "test", "The database password")
)

func main() {
	flag.Parse()

	// connect to the database
	DB, err := cms.NewDatabase(cms.DatabaseConfig{
		Dialect: *dialect,
		Source: fmt.Sprintf(
			"host=%s port=%s sslmode=disable dbname=%s user=%s password=%s",
			*databaseHost,
			*databasePort,
			*database,
			*user,
			*pass,
		),
		Debug: true,
	})
	defer DB.Close()
	if err != nil {
		panic(err)
	}

	// Business domain.
	service := cms.NewService(DB.Debug(), config.Config{
		Secret: "secret",
	})

	if err := api.Start(service); err != nil {
		log.Fatalln(err)
	}
}
