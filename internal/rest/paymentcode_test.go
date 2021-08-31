package rest_test

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/golang/mock/gomock"
	golangtraining "github.com/xtianatilano/christian-golang-training-beginner"
	"github.com/xtianatilano/christian-golang-training-beginner/internal/rest"
	"github.com/xtianatilano/christian-golang-training-beginner/internal/rest/mocks"

	"github.com/julienschmidt/httprouter"
	"github.com/stretchr/testify/require"
)

func TestCreate(t *testing.T) {
	testCases := []struct {
		desc           string
		service        *mocks.MockService
		expectedReturn *golangtraining.PaymentCode
		url            string
		body           io.Reader
		expectedCode   int
	}{
		{
			desc: "create payment codes - success",
			service: func() *mocks.MockService {
				ctrl := gomock.NewController(t)
				m := mocks.NewMockService(ctrl)

				m.
					EXPECT().
					Create(gomock.Any()).
					Return(nil)

				return m
			}(),
			expectedReturn: &golangtraining.PaymentCode{
				Name: "lechsa",
			},
			body: strings.NewReader(`
				{
					"payment_code": "test",
					"name": "lechsa"
				}
			`),
			expectedCode: http.StatusCreated,
			url:          "/payment-codes",
		},
		{
			desc: "create payment codes - success",
			service: func() *mocks.MockService {
				ctrl := gomock.NewController(t)
				m := mocks.NewMockService(ctrl)

				m.
					EXPECT().
					Create(gomock.Any()).
					Return(nil)

				return m
			}(),
			expectedReturn: &golangtraining.PaymentCode{
				Name: "lechsa",
			},
			body: strings.NewReader(`
				{
					"payment_code": "test",
					"name": "lechsa"
				}
			`),
			expectedCode: http.StatusCreated,
			url:          "/payment-codes",
		},
		{
			desc: "create payment codes - success",
			service: func() *mocks.MockService {
				ctrl := gomock.NewController(t)
				m := mocks.NewMockService(ctrl)

				m.
					EXPECT().
					Create(gomock.Any()).
					Return(nil)

				return m
			}(),
			expectedReturn: &golangtraining.PaymentCode{
				Name: "lechsa",
			},
			body: strings.NewReader(`
				{
					"payment_code": "test",
					"name": "lechsa"
				}
			`),
			expectedCode: http.StatusCreated,
			url:          "/payment-codes",
		},
		{
			desc: "create payment codes - failed because name not provided",
			service: func() *mocks.MockService {
				ctrl := gomock.NewController(t)
				m := mocks.NewMockService(ctrl)

				return m
			}(),
			expectedReturn: nil,
			body: strings.NewReader(`
				{
					"payment_code": "test",
					"name": ""
				}
			`),
			expectedCode: http.StatusBadRequest,
			url:          "/payment-codes",
		},
		{
			desc: "create payment codes - failed from service",
			service: func() *mocks.MockService {
				ctrl := gomock.NewController(t)
				m := mocks.NewMockService(ctrl)

				m.
					EXPECT().
					Create(gomock.Any()).
					Return(errors.New("internal server error"))

				return m
			}(),
			expectedReturn: nil,
			body: strings.NewReader(`
				{
					"payment_code": "test",
					"name": "lechsa"
				}
			`),
			expectedCode: http.StatusInternalServerError,
			url:          "/payment-codes",
		},
	}

	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			r := httprouter.New()
			rest.InitPaymentCodeRESTHandler(r, tC.service)

			req := httptest.NewRequest("POST", tC.url, tC.body)
			req.Header.Set("Content-Type", "application/json")
			rec := httptest.NewRecorder()

			r.ServeHTTP(rec, req)
			res, _ := ioutil.ReadAll(rec.Result().Body)
			var pRes *golangtraining.PaymentCode
			if err := json.Unmarshal(res, &pRes); err != nil {
				fmt.Println("error")
				return
			}

			require.Equal(t, tC.expectedCode, rec.Code)

			if tC.expectedCode < 400 {
				require.Equal(t, tC.expectedReturn.Name, pRes.Name)
			}

		})
	}
}

func TestGetByID(t *testing.T) {
	testCases := []struct {
		desc           string
		service        *mocks.MockService
		expectedReturn golangtraining.PaymentCode
		url            string
		expectedCode   int
	}{
		{
			desc: "get payment codes by ID - success",
			service: func() *mocks.MockService {
				ctrl := gomock.NewController(t)
				m := mocks.NewMockService(ctrl)

				p := golangtraining.PaymentCode{
					Name:        "test",
					PaymentCode: "test-1",
				}

				m.
					EXPECT().
					GetByID(gomock.Any()).
					Return(p, nil)

				return m
			}(),
			expectedReturn: golangtraining.PaymentCode{
				Name:        "test",
				PaymentCode: "test-1",
			},
			expectedCode: http.StatusOK,
			url:          "/payment-codes/id-123",
		},
		{
			desc: "get payment codes by ID - failed",
			service: func() *mocks.MockService {
				ctrl := gomock.NewController(t)
				m := mocks.NewMockService(ctrl)

				m.
					EXPECT().
					GetByID(gomock.Any()).
					Return(golangtraining.PaymentCode{}, errors.New("error from server"))

				return m
			}(),
			expectedReturn: golangtraining.PaymentCode{
				Name:        "test",
				PaymentCode: "test-1",
			},
			expectedCode: http.StatusNotFound,
			url:          "/payment-codes/id-123",
		},
	}

	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			r := httprouter.New()
			rest.InitPaymentCodeRESTHandler(r, tC.service)

			req := httptest.NewRequest("GET", tC.url, nil)
			req.Header.Set("Content-Type", "application/json")
			rec := httptest.NewRecorder()

			r.ServeHTTP(rec, req)

			require.Equal(t, tC.expectedCode, rec.Code)
		})
	}
}
