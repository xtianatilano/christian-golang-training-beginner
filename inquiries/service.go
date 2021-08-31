package inquiries

import (
	golangtraining "github.com/xtianatilano/christian-golang-training-beginner"
)

//go:generate mockgen -destination=mocks/mock_inquiries_repo.go -package=mocks . Repository
type Repository interface {
	Create(p *golangtraining.Inquiry) (*golangtraining.Inquiry, error)
	GetByPaymentCode(p string) (golangtraining.Inquiry, error)
}

//go:generate mockgen -destination=mocks/mock_payment_code_service.go -package=mocks . PaymentCodeService
type PaymentCodeService interface {
	GetByPaymentCode(paymentCode string) (res golangtraining.PaymentCode, err error)
}

type InquiryService struct {
	repo               Repository
	paymentCodeService PaymentCodeService
}

// NewService will initialize the implementations of VA Settings service
func NewService(
	repo Repository,
	paymentCodeService PaymentCodeService,
) *InquiryService {
	return &InquiryService{
		repo:               repo,
		paymentCodeService: paymentCodeService,
	}
}

func (i InquiryService) Create(p *golangtraining.Inquiry) (*golangtraining.Inquiry, error) {
	// check if payment code exist
	_, err := i.paymentCodeService.GetByPaymentCode(p.PaymentCode)
	if err != nil {
		return nil, err
	}

	res, err := i.repo.Create(p)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (i InquiryService) GetByPaymentCode(p string) (golangtraining.Inquiry, error) {
	// check if payment code exist
	res, err := i.repo.GetByPaymentCode(p)
	if err != nil {
		return res, err
	}

	return res, nil
}
