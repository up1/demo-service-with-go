package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

type helloWorldResponse struct {
	Header headerResponse `json:"header"`
	Body   bodyResponse   `json:"body"`
}
type headerResponse struct {
	Code        int    `json:"code"`
	Description string `json:"description"`
}
type bodyResponse struct {
	Message string `json:"message"`
}

type helloWorld2Response struct {
	Message string `json:"message"`
}
type helloWorld2Request struct {
	Name string `json:"name"`
}

func main() {
	http.HandleFunc("/hello", helloHandler)
	http.HandleFunc("/hello2", hello2Handler)
	http.ListenAndServe(":8080", nil)
}

func hello2Handler(w http.ResponseWriter, r *http.Request) {
	// Request
	var request helloWorld2Request
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&request)
	if err != nil {
		http.Error(w, "Bad request", http.StatusBadRequest)
		return
	}
	// Response
	response := helloWorld2Response{
		Message: fmt.Sprintf("Hello %v", request.Name)}
	encoder := json.NewEncoder(w)
	encoder.Encode(&response)
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	name, _ := os.Hostname()
	response := helloWorldResponse{
		Header: headerResponse{
			Code:        200,
			Description: "Success",
		},
		Body: bodyResponse{Message: fmt.Sprintf("Hello World %v", name)},
	}
	encoder := json.NewEncoder(w)
	encoder.Encode(&response)
}
