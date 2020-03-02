package main

import (
	"log"
	"net/http"
	"os"
	"os/signal"

	"github.com/BOPR/config"
	"github.com/BOPR/listener"
	"github.com/BOPR/types"
	"github.com/gorilla/mux"
)

func main() {
	config.ParseAndInitGlobalConfig()
	syncer := listener.NewSyncer()
	syncer.DBInstance.StoreListenerLog(types.ListenerLog{LastRecordedBlock: "101"})
	if err := syncer.Start(); err != nil {
		log.Fatalln("Unable to start syncer", "error")
	} // go routine to catch signal
	catchSignal := make(chan os.Signal, 1)
	signal.Notify(catchSignal, os.Interrupt)
	go func() {
		// sig is a ^C, handle it
		for range catchSignal {

			syncer.Stop()
			// exit
			os.Exit(1)
		}
	}()
	r := mux.NewRouter()
	http.ListenAndServe(":8080", r)

}
