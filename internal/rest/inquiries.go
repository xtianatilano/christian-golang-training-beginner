package rest

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	golangtraining "github.com/xtianatilano/christian-golang-training-beginner"
	"github.com/julienschmidt/httprouter"
)

//go:generate mockgen -destination=mocks/mock_inquiries_service.go -package=mocks . InquiryService
type InquiryService interface {
	Create(p *golangtraining.Inquiry) (*golangtraining.Inquiry, error)
}

type inquiryServiceHandler struct {
	service InquiryService
}

// InitPaymentCodeRESTHandler will initialize the REST handler for Payment Code
func InitInquiryRESTHandler(r *httprouter.Router, service InquiryService) {
	h := inquiryServiceHandler{
		service: service,
	}

	r.POST("/inquiry", h.Create)
}

func (i inquiryServiceHandler) Create(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	w.Header().Set("Content-Type", "application/json")

	b, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`{"message":"error"}`))
		return
	}

	var p golangtraining.Inquiry
	if err = json.Unmarshal(b, &p); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`{"message":"error"}`))
		return
	}

	res, err := i.service.Create(&p)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`{"message":"error creating"}`))
		return
	}

	e, err := json.Marshal(res)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`{"message":"error"}`))
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Write(e)
}
