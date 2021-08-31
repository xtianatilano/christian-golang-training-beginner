package inquiries_test

import (
	"errors"
	"testing"

	"github.com/xtianatilano/christian-golang-training-beginner/inquiries"

	"github.com/golang/mock/gomock"
	golangtraining "github.com/xtianatilano/christian-golang-training-beginner"
	"github.com/xtianatilano/christian-golang-training-beginner/inquiries/mocks"
	"github.com/stretchr/testify/require"
)

func TestCreatePaymentCode(t *testing.T) {
	testCases := []struct {
		desc               string
		repo               *mocks.MockRepository
		paymentCodeService *mocks.MockPaymentCodeService
		expectedReturn     error
	}{
		{
			desc: "create payment codes - success",
			repo: func() *mocks.MockRepository {
				ctrl := gomock.NewController(t)
				m := mocks.NewMockRepository(ctrl)

				paymentCode := &golangtraining.Inquiry{
					ID:            "1a510335-83eb-49f4-a121-9dc2d7a11348",
					PaymentCode:   "payment-code-1",
					TransactionId: "tid-123",
				}

				m.
					EXPECT().
					Create(gomock.Any()).
					Return(paymentCode, nil)

				return m
			}(),
			paymentCodeService: func() *mocks.MockPaymentCodeService {
				ctrl := gomock.NewController(t)
				m := mocks.NewMockPaymentCodeService(ctrl)

				paymentCode := golangtraining.PaymentCode{
					ID:          "1a510335-83eb-49f4-a121-9dc2d7a11348",
					PaymentCode: "payment-code-1",
				}

				m.
					EXPECT().
					GetByPaymentCode(gomock.Any()).
					Return(paymentCode, nil)

				return m
			}(),
			expectedReturn: nil,
		},
		{
			desc: "create payment codes - return error from repository",
			repo: func() *mocks.MockRepository {
				ctrl := gomock.NewController(t)
				m := mocks.NewMockRepository(ctrl)

				return m
			}(),
			paymentCodeService: func() *mocks.MockPaymentCodeService {
				ctrl := gomock.NewController(t)
				m := mocks.NewMockPaymentCodeService(ctrl)

				m.
					EXPECT().
					GetByPaymentCode(gomock.Any()).
					Return(golangtraining.PaymentCode{}, errors.New("unknown error"))

				return m
			}(),
			expectedReturn: errors.New("unknown error"),
		},
	}

	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			service := inquiries.NewService(tC.repo, tC.paymentCodeService)
			_, err := service.Create(&golangtraining.Inquiry{
				PaymentCode: "payment-code-1",
			})

			require.Equal(t, tC.expectedReturn, err)
		})
	}
}
