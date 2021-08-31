package rest

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	golangtraining "github.com/xtianatilano/christian-golang-training-beginner"
	"github.com/julienschmidt/httprouter"
)

//go:generate mockgen -destination=mocks/mock_payments_service.go -package=mocks . PaymentService
type PaymentService interface {
	Create(p *golangtraining.Payment) (*golangtraining.Payment, error)
}

type paymentServiceHandler struct {
	service PaymentService
}

// InitPaymentCodeRESTHandler will initialize the REST handler for Payment
func InitPaymentRESTHandler(r *httprouter.Router, service PaymentService) {
	h := paymentServiceHandler{
		service: service,
	}

	r.POST("/payment", h.Create)
}

func (i paymentServiceHandler) Create(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	w.Header().Set("Content-Type", "application/json")

	b, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`{"message":"error"}`))
		return
	}

	var p golangtraining.Payment
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
