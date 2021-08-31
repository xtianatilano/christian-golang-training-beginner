package golangtraining

import (
	"time"
)

type Inquiry struct {
	ID            string    `json:"id"`
	PaymentCode   string    `json:"payment_code"`
	TransactionId string    `json:"transaction_id"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
}
