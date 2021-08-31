package cmd

import (
	"fmt"
	"log"
	"time"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(expireCronCommand)
}

var expireCronCommand = &cobra.Command{
	Use:   "expire-payment-code",
	Short: "Command to start job to expire payment code",
	Run:   expireVirtualAccounts,
}

func expireVirtualAccounts(
	_ *cobra.Command,
	_ []string,
) {
	fmt.Print("running cron")
	// Manually run the Expire CRON job every 5s since the lowest granularity is minute in a CRON expression for k8s
	for i := 1; i < 12; i++ {

		err := expirePaymentCodesJob.Work()
		if err != nil {
			log.Fatal(err)
		}
		time.Sleep(5 * time.Second)
	}
}
