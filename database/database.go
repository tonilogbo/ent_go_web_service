package database

import (
	"entrds/ent"
	"fmt"
	"log"
	"os"
)

var Client *ent.Client

func SetupDatabase() {
	var err error
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbEndpoint := os.Getenv("DB_ENDPOINT")
	databaseName := os.Getenv("DB_NAME")
	dataSourceName := fmt.Sprintf(`%v:%v@tcp(%v)/%v?parseTime=true`, dbUser, dbPassword, dbEndpoint, databaseName)
	Client, err = ent.Open("mysql", dataSourceName)
	if err != nil {
		log.Fatalf("failed opening connection to database: %v", err)
	}
}