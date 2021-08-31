package cmd

import (
	"database/sql"
	"fmt"
	"log"
	"strconv"

	"github.com/aws/aws-sdk-go/aws/endpoints"

	"github.com/xtianatilano/christian-golang-training-beginner/internal/sqs"

	"github.com/xtianatilano/christian-golang-training-beginner/app/cmd/helpers"

	"github.com/spf13/cobra"
	"github.com/xtianatilano/christian-golang-training-beginner/internal/jobs"

	"github.com/xtianatilano/christian-golang-training-beginner/inquiries"
	"github.com/xtianatilano/christian-golang-training-beginner/internal/postgres"
	"github.com/xtianatilano/christian-golang-training-beginner/paymentcode"
	"github.com/xtianatilano/christian-golang-training-beginner/payments"

	_ "github.com/lib/pq"
)

var (
	paymentCodeRepository *postgres.PaymentCodeRepository
	paymentCodeService    *paymentcode.PaymentCodeService
	inquiriesRepository   *postgres.InquiriesRepository
	inquiriesService      *inquiries.InquiryService
	paymentsService       *payments.PaymentService
	paymentsRepository    *postgres.PaymentsRepository
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

	psqlInfo := fmt.Sprintf("postgres://%v:%v@%v:%v/%v?sslmode=disable",
		user,
		password,
		host,
		port,
		dbname,
	)

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	sqsPublisher, _ := sqs.NewPublisher(endpoints.UsEast1RegionID, helpers.MustHaveEnv("SQS_ENDPOINT"))

	paymentCodeRepository = postgres.NewPaymentCodeRepository(db)
	paymentCodeService = paymentcode.NewService(paymentCodeRepository)
	inquiriesRepository = postgres.NewInquiriesRepository(db)
	inquiriesService = inquiries.NewService(inquiriesRepository, *paymentCodeService)
	paymentsRepository = postgres.NewPaymentsRepository(db)
	paymentsService = payments.NewService(paymentsRepository, *inquiriesService, sqsPublisher)

	fmt.Println("Successfully connected!")

	expirePaymentCodesJob = jobs.ExpirePaymentCodesJob{PaymentCodesService: paymentCodeService}
}

// Execute will call the root command execute
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		log.Fatal(err)
	}
}