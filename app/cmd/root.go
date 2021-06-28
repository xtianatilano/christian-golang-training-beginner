package cmd

import (
	"log"
	"fmt"
	"database/sql"
	_ "github.com/lib/pq"

	"github.com/xtianatilano/christian-golang-training-beginner/internal/postgres"
	"github.com/xtianatilano/christian-golang-training-beginner/paymentcode"
)

var (
	paymentCodeRepository *postgres.Repository
	service               *paymentcodeservice.Service
)

func init() {
	dataSourceName := "postgres://commander:pass123$$@localhost:5432/xendit?sslmode=disable"

	db, err := sql.Open("postgres", dataSourceName)
	if err != nil {
		log.Fatal(err)
	}
	paymentCodeRepository := postgres.New(db)
	paymentcodeservice.New(paymentCodeRepository)

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	fmt.Println("Successfully connected!")
}
