package postgres

import (
	"database/sql"
	"time"

	sq "github.com/Masterminds/squirrel"
	"github.com/google/uuid"
	golangtraining "github.com/xtianatilano/christian-golang-training-beginner"
	"github.com/pkg/errors"
)

type InquiriesRepository struct {
	DB *sql.DB
}

func NewInquiriesRepository(db *sql.DB) *InquiriesRepository {
	return &InquiriesRepository{
		DB: db,
	}
}

func (t InquiriesRepository) Create(p *golangtraining.Inquiry) (*golangtraining.Inquiry, error) {
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
		Insert("inquiries").
		Columns("id", "payment_code", "transaction_id", "created_at", "updated_at").
		Values(p.ID, p.PaymentCode, p.TransactionId, p.CreatedAt, p.UpdatedAt).
		Suffix("RETURNING \"id\"").
		PlaceholderFormat(sq.Dollar)

	_, err = query.RunWith(t.DB).Exec()
	if err != nil {
		err = errors.Wrap(err, "error creating data")
		return nil, err
	}

	return p, nil
}

func (t InquiriesRepository) GetByPaymentCode(p string) (golangtraining.Inquiry, error) {
	var res golangtraining.Inquiry
	var err error

	query := sq.
		Select("*").
		Where(sq.Eq{"payment_code": p}).
		From("inquiries").
		PlaceholderFormat(sq.Dollar)

	err = query.RunWith(t.DB).QueryRow().Scan(
		&res.ID, &res.PaymentCode, &res.TransactionId, &res.CreatedAt, &res.UpdatedAt,
	)

	if err != nil {
		return res, errors.New("not found")
	}

	return res, nil
}
