package main

import (
	"encoding/json"
	"fmt"
	"log"
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

func notFoundHandler(w http.ResponseWriter, r *http.Request) {
	jsonData := &defaultResponse{ Message: "Route does not exist!" }
	handleDefaultResponse(w, http.StatusNotFound, jsonData)
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

func handleRequests() {
	port := "9000"
	server := http.Server{
		Addr:    "0.0.0.0:" + port,
		Handler: nil,
	}
	http.HandleFunc("/", notFoundHandler)
	http.HandleFunc("/health", healthCheck)
	http.HandleFunc("/hello-world", helloWorld)

	log.Println("Listening on port:", port)
	log.Fatal(server.ListenAndServe())
}

func main() {
	handleRequests()
}