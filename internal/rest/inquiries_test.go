package rest_test

import (
	"encoding/json"
	"fmt"
	"github.com/pkg/errors"
	"io"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/xtianatilano/christian-golang-training-beginner/internal/rest"

	"github.com/golang/mock/gomock"
	golangtraining "github.com/xtianatilano/christian-golang-training-beginner"
	"github.com/xtianatilano/christian-golang-training-beginner/internal/rest/mocks"

	"github.com/julienschmidt/httprouter"
	"github.com/stretchr/testify/require"
)

func TestCreateInquiries(t *testing.T) {
	testCases := []struct {
		desc           string
		service        *mocks.MockInquiryService
		expectedReturn *golangtraining.Inquiry
		url            string
		body           io.Reader
		expectedCode   int
	}{
		{
			desc: "create inquiries - success",
			service: func() *mocks.MockInquiryService {
				ctrl := gomock.NewController(t)
				m := mocks.NewMockInquiryService(ctrl)

				p := &golangtraining.Inquiry{
					ID:            "c2636568-057e-45a6-a23b-a2fecf67376d",
					PaymentCode:   "hello",
					TransactionId: "1234",
				}

				m.
					EXPECT().
					Create(gomock.Any()).
					Return(p, nil)

				return m
			}(),
			expectedReturn: &golangtraining.Inquiry{
				ID:            "c2636568-057e-45a6-a23b-a2fecf67376d",
				PaymentCode:   "hello",
				TransactionId: "1234",
			},
			body: strings.NewReader(`
				{
					"transaction_id": "1234",
					"payment_code": "hello"
				}
			`),
			expectedCode: http.StatusCreated,
			url:          "/inquiry",
		},
		{
			desc: "create inquiries - failed",
			service: func() *mocks.MockInquiryService {
				ctrl := gomock.NewController(t)
				m := mocks.NewMockInquiryService(ctrl)



				m.
					EXPECT().
					Create(gomock.Any()).
					Return(nil, errors.New("internal server error"))

				return m
			}(),
			expectedReturn: &golangtraining.Inquiry{
				ID:            "c2636568-057e-45a6-a23b-a2fecf67376d",
				PaymentCode:   "hello",
				TransactionId: "1234",
			},
			body: strings.NewReader(`
				{
					"transaction_id": "1234",
					"payment_code": "hello"
				}
			`),
			expectedCode: http.StatusInternalServerError,
			url:          "/inquiry",
		},
	}

	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			r := httprouter.New()
			rest.InitInquiryRESTHandler(r, tC.service)

			req := httptest.NewRequest("POST", tC.url, tC.body)
			req.Header.Set("Content-Type", "application/json")
			rec := httptest.NewRecorder()

			r.ServeHTTP(rec, req)
			res, _ := ioutil.ReadAll(rec.Result().Body)
			var pRes *golangtraining.Inquiry
			if err := json.Unmarshal(res, &pRes); err != nil {
				fmt.Println("error")
				return
			}

			require.Equal(t, tC.expectedCode, rec.Code)

			if tC.expectedCode < 400 {
				require.Equal(t, tC.expectedReturn.PaymentCode, pRes.PaymentCode)
				require.Equal(t, tC.expectedReturn.TransactionId, pRes.TransactionId)
			}

		})
	}
}
