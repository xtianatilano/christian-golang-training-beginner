package postgres

import (
	"database/sql"
	"time"

	"github.com/google/uuid"
	"github.com/xtianatilano/christian-golang-training-beginner"
	"github.com/xtianatilano/christian-golang-training-beginner/pkg"
)

type Repository struct {
	db *sql.DB
}

func New(db *sql.DB) Repository {
	return Repository{
		db: db,
	}
}

func (r Repository) Create(paymentCode *paymentcodedomain.PaymentCode) (standardError *standarderror.StandardError) {
	paymentCode.Id = uuid.NewString()
	paymentCode.CreatedAt = time.Now().UTC()
	paymentCode.UpdatedAt = time.Now().UTC()
	sqlStatement := `INSERT INTO paymentcode (id, payment_code, name, status, expiration_date, 
		created_at, updated_at) VALUES ($1, $2, $3, $4, $5, $6, $7)`

	result, err := r.db.Exec(
		sqlStatement,
		paymentCode.Id,
		paymentCode.PaymentCode,
		paymentCode.Name,
		paymentCode.Status,
		paymentCode.ExpirationDate,
		paymentCode.CreatedAt,
		paymentCode.UpdatedAt,
	)

	insertedRow, err := result.RowsAffected()

	if insertedRow != 1 {
		standardError = &standarderror.StandardError{
			ErrorCode:    "CREATE_PAYMENT_CODE_ERROR",
			ErrorMessage: "fail to create payment code",
			StatusCode:   500,
		}
		return
	}
	if err != nil {
		standardError = &standarderror.StandardError{
			ErrorCode:    "CREATE_PAYMENT_CODE_ERROR",
			ErrorMessage: "database error",
			StatusCode:   500,
		}
		return
	}
	return nil
}

func (r Repository) Get(id string) (standardError *standarderror.StandardError, paymentCode paymentcodedomain.PaymentCode) {
	sqlStatement :=
		`SELECT
	id,
	payment_code,
	name, status,
	expiration_date,
	created_at, updated_at
	FROM paymentcode WHERE id=$1`

	row := r.db.QueryRow(
		sqlStatement,
		id,
	)

	err := row.Scan(
		&paymentCode.Id,
		&paymentCode.PaymentCode,
		&paymentCode.Name,
		&paymentCode.Status,
		&paymentCode.ExpirationDate,
		&paymentCode.CreatedAt,
		&paymentCode.UpdatedAt,
	)

	if err != nil {
		standardError = &standarderror.StandardError{
			ErrorCode:    "PAYMENT_NOT_FOUND_ERROR",
			ErrorMessage: "payment code is not found",
			StatusCode:   404,
		}
		return standardError, paymentCode
	}

	return nil, paymentCode
}