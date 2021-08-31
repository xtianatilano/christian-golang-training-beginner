package rest_test

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/pkg/errors"

	"github.com/golang/mock/gomock"
	golangtraining "github.com/xtianatilano/christian-golang-training-beginner"
	"github.com/xtianatilano/christian-golang-training-beginner/internal/rest"
	"github.com/xtianatilano/christian-golang-training-beginner/internal/rest/mocks"

	"github.com/julienschmidt/httprouter"
	"github.com/stretchr/testify/require"
)

func TestCreatePayment(t *testing.T) {
	testCases := []struct {
		desc           string
		service        *mocks.MockPaymentService
		expectedReturn *golangtraining.Payment
		url            string
		body           io.Reader
		expectedCode   int
	}{
		{
			desc: "create payment - success",
			service: func() *mocks.MockPaymentService {
				ctrl := gomock.NewController(t)
				m := mocks.NewMockPaymentService(ctrl)

				p := &golangtraining.Payment{
					ID:            "c2636568-057e-45a6-a23b-a2fecf67376d",
					PaymentCode:   "hello",
					TransactionId: "1234",
					Amount:        10000,
					Name:          "Lechsa",
				}

				m.
					EXPECT().
					Create(gomock.Any()).
					Return(p, nil)

				return m
			}(),
			expectedReturn: &golangtraining.Payment{
				ID:            "c2636568-057e-45a6-a23b-a2fecf67376d",
				PaymentCode:   "hello",
				TransactionId: "1234",
				Amount:        10000,
				Name:          "Lechsa",
			},
			body: strings.NewReader(`
				{
					"transaction_id": "1234",
					"amount": 10000, 
					"name": "Lechsa", 
					"payment_code": "hello"
				}
			`),
			expectedCode: http.StatusCreated,
			url:          "/payment",
		},
		{
			desc: "create payment - failed",
			service: func() *mocks.MockPaymentService {
				ctrl := gomock.NewController(t)
				m := mocks.NewMockPaymentService(ctrl)

				m.
					EXPECT().
					Create(gomock.Any()).
					Return(nil, errors.New("internal server error"))

				return m
			}(),
			expectedReturn: &golangtraining.Payment{
				ID:            "c2636568-057e-45a6-a23b-a2fecf67376d",
				PaymentCode:   "hello",
				TransactionId: "1234",
				Amount:        10000,
				Name:          "Lechsa",
			},
			body: strings.NewReader(`
				{
					"transaction_id": "1234",
					"amount": 10000, 
					"name": "Lechsa", 
					"payment_code": "hello"
				}
			`),
			expectedCode: http.StatusInternalServerError,
			url:          "/payment",
		},
	}

	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			r := httprouter.New()
			rest.InitPaymentRESTHandler(r, tC.service)

			req := httptest.NewRequest("POST", tC.url, tC.body)
			req.Header.Set("Content-Type", "application/json")
			rec := httptest.NewRecorder()

			r.ServeHTTP(rec, req)
			res, _ := ioutil.ReadAll(rec.Result().Body)
			var pRes *golangtraining.Payment
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
