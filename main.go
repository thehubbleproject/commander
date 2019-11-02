package main

import (
	"net/http"

	"github.com/BOPR/rest"
	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/tx", rest.TxReceiverHandler).Methods("POST")
	http.Handle("/", r)
	http.ListenAndServe(":8080", r)
}
