package rest

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type HealthResponse struct {
	Status string `json:"status"`
}

type HelloResponse struct {
	Message string `json:"message"`
}

func healthCheck(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	w.Header().Set("Content-Type", "application/json")

	w.WriteHeader(http.StatusOK)
	jsonData := &HealthResponse{Status: "healthy"}

	e, err := json.Marshal(jsonData)

	if err != nil {
		fmt.Println(err)
		return
	}

	w.Write(e)
}

func helloWorld(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	w.Header().Set("Content-Type", "application/json")

	w.WriteHeader(http.StatusOK)
	jsonData := &HelloResponse{Message: "hello world"}

	e, err := json.Marshal(jsonData)

	if err != nil {
		fmt.Println(err)
		return
	}

	w.Write(e)
}

func InitHandler(r *httprouter.Router) {

	r.GET("/health", healthCheck)
	r.GET("/hello-world", helloWorld)
}
