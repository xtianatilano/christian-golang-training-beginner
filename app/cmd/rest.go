package cmd

import (
	"github.com/spf13/cobra"
	"log"
	"net/http"
	"time"

	"github.com/julienschmidt/httprouter"
	rest "github.com/xtianatilano/christian-golang-training-beginner/internal/rest"
)

var restCommand = &cobra.Command{
	Use:   "rest",
	Short: "Start REST server",
	Run:   restServer,
}

func init() {
	rootCmd.AddCommand(restCommand)
}

func restServer(cmd *cobra.Command, args []string) {
	port := ":3000"

	r := httprouter.New()

	initHttpClient()
	rest.InitHandler(r)

	rest.InitPaymentCodeRESTHandler(r, paymentCodeService)

	log.Println("listen on", port)
	log.Fatal(http.ListenAndServe(port, r))
}

func initHttpClient() *http.Client {
	return &http.Client{
		Transport: &http.Transport{
			MaxIdleConnsPerHost: 20,
		},
		Timeout: 10 * time.Second,
	}
}
