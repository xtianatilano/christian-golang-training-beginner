package jobs

type PaymentCodesService interface {
	Expire() error
}

// ExpireVirtualAccountsJob expires Virtual Accounts via a CRON
type ExpirePaymentCodesJob struct {
	PaymentCodesService PaymentCodesService
}

// Work is the main function of the job
func (j ExpirePaymentCodesJob) Work() (err error) {
	err = j.PaymentCodesService.Expire()
	if err != nil {
		return err
	}

	return
}
