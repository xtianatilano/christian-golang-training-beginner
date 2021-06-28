package cmd

import (
	"database/sql"
	"github.com/xtianatilano/christian-golang-training-beginner/internal/postgres"
	paymentcodeservice "github.com/xtianatilano/christian-golang-training-beginner/paymentcode"
	"log"
	"net/http"

	"github.com/xtianatilano/christian-golang-training-beginner/internal/rest"
)

func Execute() {
	dataSourceName := "postgres://commander:pass123$$@localhost:5432/xendit?sslmode=disable"

	db, err := sql.Open("postgres", dataSourceName)
	if err != nil {
		log.Fatal(err)
	}
	repo:= postgres.New(db)
	service := paymentcodeservice.New(repo)

	port := ":9000"
	server := http.NewServeMux()

	rest.HandleRequests(server)
	rest.HandlePaymentCodeRequest(server, service)

	log.Println("Listening on port", port)
	log.Fatal(http.ListenAndServe(port, server))
}