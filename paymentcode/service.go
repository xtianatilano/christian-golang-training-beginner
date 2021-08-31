package paymentcode

import (
	"fmt"
	"time"

	golangtraining "github.com/xtianatilano/christian-golang-training-beginner"
)

type GetByIDResponse struct {
	ID string
}

//go:generate mockgen -destination=mocks/mock_paymentcodes_repo.go -package=mocks . Repository
type Repository interface {
	Create(p *golangtraining.PaymentCode) (*golangtraining.PaymentCode, error)
	GetByID(ID string) (golangtraining.PaymentCode, error)
	Expire() error
}

type PaymentCodeService struct {
	repo Repository
}

// NewService will initialize the implementations of VA Settings service
func NewService(
	repo Repository,
) *PaymentCodeService {
	return &PaymentCodeService{
		repo: repo,
	}
}

func (s PaymentCodeService) Create(p *golangtraining.PaymentCode) error {
	now := time.Now().UTC()
	p.ExpirationDate = now.AddDate(51, 0, 0)

	_, err := s.repo.Create(p)
	if err != nil {
		return err
	}
	return nil
}

func (s PaymentCodeService) GetByID(ID string) (res golangtraining.PaymentCode, err error) {
	res, err = s.repo.GetByID(ID)
	if err != nil {
		fmt.Println(err)
		return
	}

	return
}

func (s PaymentCodeService) Expire() error {
	err := s.repo.Expire()
	if err != nil {
		return err
	}

	return nil
}
