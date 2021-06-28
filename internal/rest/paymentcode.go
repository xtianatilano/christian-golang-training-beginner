package rest

import (
	"encoding/json"
	_ "github.com/lib/pq"
	"log"
	"net/http"
	"regexp"
	"strings"

	paymentcodedomain "github.com/xtianatilano/christian-golang-training-beginner"
	standarderror "github.com/xtianatilano/christian-golang-training-beginner/pkg"
)

type IPaymentCodeService interface {
	Create(paymentcode *paymentcodedomain.PaymentCode) *standarderror.StandardError
	Get(id string) (err *standarderror.StandardError, paymentcode paymentcodedomain.PaymentCode)
}

type paymentCodeHandler struct {
	service IPaymentCodeService
}

func HandlePaymentCodeRequest(router *http.ServeMux, service IPaymentCodeService) {
	h := paymentCodeHandler{
		service: service,
	}
	router.HandleFunc("/", h.handleHttpMethod)
}

func (p paymentCodeHandler) handleHttpMethod(w http.ResponseWriter, r *http.Request) {
	pattern := `^/payment-codes.*`
	url := r.URL.Path
	matched, err := regexp.MatchString(pattern, r.URL.Path)
	if err != nil {
		log.Fatal(err)
	}

	if !matched {
		http.NotFound(w, r)
		return
	}

	id := strings.Replace(url, "/payment-codes/", "", 1)

	switch r.Method {
	case http.MethodPost:
		p.Create(w, r)
	case http.MethodGet:
		p.Get(w, r, id)
	default:
		http.NotFound(w, r)
	}
}

func (p paymentCodeHandler) Create(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	paymentcode := paymentcodedomain.PaymentCode{}
	json.NewDecoder(r.Body).Decode(&paymentcode)

	err := p.service.Create(&paymentcode)

	if err != nil {
		errorHandler(w, r, err)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(paymentcode)
}

func (p paymentCodeHandler) Get(w http.ResponseWriter, r *http.Request, id string) {
	w.Header().Add("Content-Type", "application/json")
	err, paymentCode := p.service.Get(id)
	if err != nil {
		errorHandler(w, r, err)
		return
	}
	json.NewEncoder(w).Encode(paymentCode)
}

func errorHandler(w http.ResponseWriter, r *http.Request, err *standarderror.StandardError) {
	errResponse := map[string]string{}
	errResponse["error_code"] = err.ErrorCode
	errResponse["message"] = err.ErrorMessage
	w.WriteHeader(err.StatusCode)

	json.NewEncoder(w).Encode(errResponse)
}