package cmd

import (
	"database/sql"
	"fmt"
	"github.com/xtianatilano/christian-golang-training-beginner/app/cmd/helpers"
	"log"
	"strconv"

	"github.com/spf13/cobra"
	"github.com/xtianatilano/christian-golang-training-beginner/internal/jobs"

	_ "github.com/lib/pq"
	"github.com/xtianatilano/christian-golang-training-beginner/internal/postgres"
	paymentcode "github.com/xtianatilano/christian-golang-training-beginner/paymentcode"
)

var (
	paymentCodeRepository *postgres.PaymentCodeRepository
	paymentCodeService    *paymentcode.PaymentCodeService
	rootCmd               = &cobra.Command{
		Use:   "app",
		Short: "Application",
	}
	expirePaymentCodesJob jobs.ExpirePaymentCodesJob
)

func init() {
	cobra.OnInitialize(initApp)
}

func initApp() {
	host := helpers.MustHaveEnv("POSTGRES_HOST")
	portStr := helpers.MustHaveEnv("POSTGRES_PORT")
	port, err := strconv.Atoi(portStr)
	if err != nil {
		log.Fatal(err, "POSTGRES_PORT is not well set ")
	}
	user := helpers.MustHaveEnv("POSTGRES_USER")
	password := helpers.MustHaveEnv("POSTGRES_PASSWORD")
	dbname := helpers.MustHaveEnv("POSTGRES_DB_NAME")


	//psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
	//	"password=%s dbname=%s sslmode=disable",
	//	host, port, user, password, dbname)
	psqlInfo := fmt.Sprintf("postgres://%v:%v@%v:%v/%v?sslmode=disable",
		user,
		password,
		host,
		port,
		dbname,
	)

	db, err := sql.Open("postgres", psqlInfo)
	paymentCodeRepository = postgres.NewPaymentCodeRepository(db)
	paymentCodeService = paymentcode.NewService(paymentCodeRepository)

	if err != nil {
		panic(err)
	}

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	fmt.Println("Successfully connected!")

	expirePaymentCodesJob = jobs.ExpirePaymentCodesJob{PaymentCodesService: paymentCodeService}
}

// Execute will call the root command execute
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		log.Fatal(err)
	}
}
