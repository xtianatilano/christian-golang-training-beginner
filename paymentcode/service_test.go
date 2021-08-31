package paymentcode_test

import (
	"errors"
	"github.com/xtianatilano/christian-golang-training-beginner/paymentcode"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
	golangtraining "github.com/xtianatilano/christian-golang-training-beginner"
	"github.com/xtianatilano/christian-golang-training-beginner/paymentcode/mocks"
)

func TestCreatePaymentCode(t *testing.T) {
	testCases := []struct {
		desc           string
		repo           *mocks.MockRepository
		expectedReturn error
	}{
		{
			desc: "create payment codes - success",
			repo: func() *mocks.MockRepository {
				ctrl := gomock.NewController(t)
				m := mocks.NewMockRepository(ctrl)

				paymentCode := &golangtraining.PaymentCode{
					ID:          "1a510335-83eb-49f4-a121-9dc2d7a11348",
					PaymentCode: "payment-code-1",
					Name:        "Name",
					Status:      "ACTIVE",
				}

				m.
					EXPECT().
					Create(gomock.Any()).
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

				m.
					EXPECT().
					Create(gomock.Any()).
					Return(nil, errors.New("unknown error"))

				return m
			}(),

			expectedReturn: errors.New("unknown error"),
		},
	}

	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			service := paymentcode.NewService(tC.repo)
			err := service.Create(&golangtraining.PaymentCode{})

			require.Equal(t, tC.expectedReturn, err)
		})
	}
}

func TestGetByID(t *testing.T) {
	testCases := []struct {
		desc           string
		repo           *mocks.MockRepository
		expectedReturn golangtraining.PaymentCode
		expectedError  error
	}{
		{
			desc: "get payment codes by ID - success",
			repo: func() *mocks.MockRepository {
				ctrl := gomock.NewController(t)
				m := mocks.NewMockRepository(ctrl)

				paymentCode := golangtraining.PaymentCode{
					ID:          "1a510335-83eb-49f4-a121-9dc2d7a11348",
					Name:        "Name",
					Status:      "ACTIVE",
					PaymentCode: "payment-code-1",
				}

				m.
					EXPECT().
					GetByID(gomock.Any()).
					Return(paymentCode, nil)

				return m
			}(),

			expectedReturn: golangtraining.PaymentCode{
				ID:          "1a510335-83eb-49f4-a121-9dc2d7a11348",
				Name:        "Name",
				Status:      "ACTIVE",
				PaymentCode: "payment-code-1",
			},
			expectedError: nil,
		},
		{
			desc: "get payment codes by ID - return error from repository",
			repo: func() *mocks.MockRepository {
				ctrl := gomock.NewController(t)
				m := mocks.NewMockRepository(ctrl)

				m.
					EXPECT().
					GetByID(gomock.Any()).
					Return(golangtraining.PaymentCode{}, errors.New("unknown error"))

				return m
			}(),

			expectedReturn: golangtraining.PaymentCode{},
			expectedError:  errors.New("unknown error"),
		},
	}

	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			service := paymentcode.NewService(tC.repo)
			res, err := service.GetByID("id")

			if err != nil {
				require.Equal(t, tC.expectedError, err)
			}

			require.Equal(t, tC.expectedReturn, res)
		})
	}
}

func TestExpire(t *testing.T) {
	testCases := []struct {
		desc           string
		repo           *mocks.MockRepository
		expectedError  error
	}{
		{
			desc: "expire payment codes - success",
			repo: func() *mocks.MockRepository {
				ctrl := gomock.NewController(t)
				m := mocks.NewMockRepository(ctrl)
				m.
					EXPECT().
					Expire().
					Return(nil)

				return m
			}(),
			expectedError: nil,
		},
		{
			desc: "expire payment codes - failed",
			repo: func() *mocks.MockRepository {
				ctrl := gomock.NewController(t)
				m := mocks.NewMockRepository(ctrl)

				m.
					EXPECT().
					Expire().
					Return(errors.New("unknown error"))

				return m
			}(),

			expectedError:  errors.New("unknown error"),
		},
	}

	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			service := paymentcode.NewService(tC.repo)
			err := service.Expire()

			require.Equal(t, tC.expectedError, err)
		})
	}
}
