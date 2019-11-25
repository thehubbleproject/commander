package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"

	db "github.com/BOPR/db"
	"github.com/BOPR/poller"
	"github.com/BOPR/rest"
	"github.com/BOPR/types"

	"github.com/gorilla/mux"
)

func main() {
	// create mgo session
	session, err := db.NewSession("mongodb://localhost:27017")
	if err != nil {
		log.Fatalf("Unable to connect to DB")
	}

	db.MgoSession = *session
	aggregator := poller.NewAggregator()
	types.ContractCallerObj, err = types.NewContractCaller()
	if err != nil {
		panic(err)
	}

	// go routine to catch signal
	catchSignal := make(chan os.Signal, 1)
	signal.Notify(catchSignal, os.Interrupt)
	go func() {
		// sig is a ^C, handle it
		for range catchSignal {
			aggregator.Stop()

			// exit
			os.Exit(1)
		}
	}()

	r := mux.NewRouter()
	r.HandleFunc("/tx", rest.TxReceiverHandler).Methods("POST")
	http.Handle("/", r)

	if err := aggregator.Start(); err != nil {
		log.Fatalln("Unable to start aggregator", "error", err)
	}
	// start server
	fmt.Printf("Starting server")
	http.ListenAndServe(":8080", r)

}
