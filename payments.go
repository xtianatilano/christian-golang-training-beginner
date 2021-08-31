package golangtraining

import (
	"time"
)

type Payment struct {
	ID            string    `json:"id"`
	PaymentCode   string    `json:"payment_code"`
	TransactionId string    `json:"transaction_id"`
	Amount        int       `json:"amount"`
	Name          string    `json:"name"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
}
