package rest

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/julienschmidt/httprouter"
	golangtraining "github.com/xtianatilano/christian-golang-training-beginner"
)

//go:generate mockgen -destination=mocks/mock_paymentcodes_service.go -package=mocks . Service
type Service interface {
	Create(p *golangtraining.PaymentCode) error
	GetByID(ID string) (golangtraining.PaymentCode, error)
}

type paymentCodeServiceHandler struct {
	service        Service
}

type GetByIDRes struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

// InitPaymentCodeRESTHandler will initialize the REST handler for Payment Code
func InitPaymentCodeRESTHandler(r *httprouter.Router, service Service) {
	h := paymentCodeServiceHandler{
		service:        service,
	}

	r.POST("/payment-codes", h.Create)
	r.GET("/payment-codes/:id", h.GetByID)
}

func (s paymentCodeServiceHandler) Create(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	w.Header().Set("Content-Type", "application/json")

	b, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`{"message":"error"}`))
		return
	}

	var p *golangtraining.PaymentCode
	if err = json.Unmarshal(b, &p); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`{"message":"error"}`))
		return
	}

	if p.Name == "" {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(`{"message":"bad request"}`))
		return
	}

	err = s.service.Create(p)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`{"message":"error creating"}`))
		return
	}

	e, err := json.Marshal(p)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`{"message":"error"}`))
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Write(e)
}

func (s paymentCodeServiceHandler) GetByID(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	w.Header().Set("Content-Type", "application/json")
	pID := ps.ByName("id")

	res, err := s.service.GetByID(pID)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte(`{"message":"not found"}`))
		return
	}

	e, err := json.Marshal(res)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`{"message":"error"}`))

		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(e)
}
