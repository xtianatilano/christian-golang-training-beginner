package paymentcodeservice

import (
	"time"

	"github.com/xtianatilano/christian-golang-training-beginner"
	"github.com/xtianatilano/christian-golang-training-beginner/pkg"
)

type IPaymentCodeRepository interface {
	Create(paymentCode *paymentcodedomain.PaymentCode) *standarderror.StandardError
	Update(paymentCode *paymentcodedomain.PaymentCode) (err *standarderror.StandardError)
	Get(id string) (err *standarderror.StandardError, paymentCode paymentcodedomain.PaymentCode)
	Delete(id string) (err *standarderror.StandardError)
}

type Service struct {
	repo IPaymentCodeRepository
}

func New(repo IPaymentCodeRepository) Service {
	return Service{
		repo: repo,
	}
}

func (s Service) Create(paymentCode *paymentcodedomain.PaymentCode) (standardError *standarderror.StandardError) {
	expirationDate := time.Now().UTC().AddDate(69, 0, 0)
	paymentCode.ExpirationDate = &expirationDate
	paymentCode.Status = paymentcodedomain.PaymentCodeStatusActive

	if paymentCode.Name == "" {
		standardError = &standarderror.StandardError{
			ErrorCode: "VALIDATION_ERROR",
			ErrorMessage: "name is is required",
			StatusCode: 400,
		}
		return
	}

	if paymentCode.PaymentCode == "" {
		standardError = &standarderror.StandardError{
			ErrorCode: "VALIDATION_ERROR",
			ErrorMessage: "payment_code is required",
			StatusCode: 400,
		}
		return
	}

	standardError = s.repo.Create(paymentCode)
	return
}

func (s Service) Update(paymentCode *paymentcodedomain.PaymentCode, id string) (err *standarderror.StandardError) {
	err, oldPaymentCode := s.repo.Get(id)

	if err != nil {
		return err
	}

	paymentCode.Id = id

	if paymentCode.Name == "" {
		paymentCode.Name = oldPaymentCode.Name
	}

	if paymentCode.PaymentCode == "" {
		paymentCode.PaymentCode = oldPaymentCode.PaymentCode
	}

	if paymentCode.ExpirationDate != nil {
		timeNow := time.Now().UTC()
		if timeNow.After(*(paymentCode.ExpirationDate)) {
			paymentCode.Status = paymentcodedomain.PaymentCodeStatusExpired
		} else {
			paymentCode.Status = paymentcodedomain.PaymentCodeStatusActive
		}
	}

	if paymentCode.ExpirationDate == nil {
		paymentCode.ExpirationDate = oldPaymentCode.ExpirationDate
	}

	err = s.repo.Update(paymentCode)
	return
}

func (s Service) Get(id string) (err *standarderror.StandardError, paymentCode paymentcodedomain.PaymentCode) {
	err, paymentCode = s.repo.Get(id)
	if err != nil {
		return err, paymentCode
	}
	return nil, paymentCode
}

func (s Service) Delete(id string) (err *standarderror.StandardError) {
	err,_ = s.repo.Get(id)
	if err != nil {
		return err
	}
	return s.repo.Delete(id)
}
