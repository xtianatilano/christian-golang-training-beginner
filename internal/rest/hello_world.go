package rest

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type healthCheckResponse struct {
	Status string `json:"status"`
}

type defaultResponse struct {
	Message string `json:"message"`
}

func healthCheck(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	jsonData := &healthCheckResponse{ Status: "healthy" }
	e, err := json.Marshal(jsonData)

	if err != nil {
		fmt.Println(err)
		return
	}
	w.Write(e)
}

func helloWorld(w http.ResponseWriter, r *http.Request){
	jsonData := &defaultResponse{ Message: "hello world" }
	handleDefaultResponse(w, http.StatusOK, jsonData)
}

func handleDefaultResponse(w http.ResponseWriter, status int, jsonData *defaultResponse) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	e, err := json.Marshal(jsonData)

	if err != nil {
		fmt.Println(err)
		return
	}

	w.Write(e)
}

func HandleRequests(router *http.ServeMux) {
	router.HandleFunc("/health", healthCheck)
	router.HandleFunc("/hello-world", helloWorld)
}