package postgres

import (
	"database/sql"
	"time"

	sq "github.com/Masterminds/squirrel"
	"github.com/google/uuid"
	golangtraining "github.com/xtianatilano/christian-golang-training-beginner"
	"github.com/pkg/errors"
)

type PaymentsRepository struct {
	DB *sql.DB
}

func NewPaymentsRepository(db *sql.DB) *PaymentsRepository {
	return &PaymentsRepository{
		DB: db,
	}
}

func (t PaymentsRepository) Create(p *golangtraining.Payment) (*golangtraining.Payment, error) {
	newUUID, err := uuid.NewRandom()
	if err != nil {
		err = errors.Wrap(err, "can't generate the UUID")
		return nil, err
	}

	p.ID = newUUID.String()
	now := time.Now().UTC()
	p.CreatedAt = now
	p.UpdatedAt = now

	query := sq.
		Insert("payments").
		Columns("id", "payment_code", "name", "amount", "transaction_id", "created_at", "updated_at").
		Values(p.ID, p.PaymentCode, p.Name, p.Amount, p.TransactionId, p.CreatedAt, p.UpdatedAt).
		Suffix("RETURNING \"id\"").
		PlaceholderFormat(sq.Dollar)

	_, err = query.RunWith(t.DB).Exec()
	if err != nil {
		err = errors.Wrap(err, "error creating data")
		return nil, err
	}

	return p, nil
}
