package payments

import (
	"fmt"

	golangtraining "github.com/xtianatilano/christian-golang-training-beginner"
	"github.com/xtianatilano/christian-golang-training-beginner/inquiries"
)

//go:generate mockgen -destination=mocks/mock_payments_repo.go -package=mocks . Repository
type Repository interface {
	Create(p *golangtraining.Payment) (*golangtraining.Payment, error)
}

//go:generate mockgen -destination=mocks/mock_publisher.go -package=mocks . Publisher
type Publisher interface {
	Publish(interface{}) error
}

type PaymentService struct {
	repo             Repository
	inquiriesService inquiries.InquiryService
	publisher        Publisher
}

// NewService will initialize the implementations of VA Settings service
func NewService(
	repo Repository,
	inquiriesService inquiries.InquiryService,
	publisher Publisher,
) *PaymentService {
	return &PaymentService{
		repo:             repo,
		inquiriesService: inquiriesService,
		publisher:        publisher,
	}
}

func (i PaymentService) Create(p *golangtraining.Payment) (*golangtraining.Payment, error) {
	// check if payment code exist
	_, err := i.inquiriesService.GetByPaymentCode(p.PaymentCode)
	if err != nil {
		return nil, err
	}

	res, err := i.repo.Create(p)
	if err != nil {
		return nil, err
	}
	err = i.publisher.Publish(p)
	if err != nil {
		fmt.Println(err)
	}

	return res, nil
}
