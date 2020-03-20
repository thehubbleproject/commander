package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"

	"github.com/BOPR/common"
	"github.com/BOPR/config"
	db "github.com/BOPR/db"
	"github.com/BOPR/listener"
	"github.com/BOPR/poller"
	"github.com/BOPR/rest"
	"github.com/BOPR/types"
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// StartCmd starts the daemon
func StartCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "start",
		Short: "Starts BOPR daemon",
		Run: func(cmd *cobra.Command, args []string) {
			// create viper object
			viperObj := viper.New()

			// get current directory
			dir, err := os.Getwd()
			common.PanicIfError(err)

			// set config paths
			viperObj.SetConfigName(ConfigFileName) // name of config file (without extension)
			viperObj.AddConfigPath(dir)

			// finally! read config
			err = viperObj.ReadInConfig()
			common.PanicIfError(err)

			// unmarshall to the configration object
			var cfg config.Configuration
			if err = viperObj.UnmarshalExact(&cfg); err != nil {
				common.PanicIfError(err)
			}
			//
			// Init global config objects
			//

			// init global config
			config.GlobalCfg = cfg
			// TODO use a better way to handle priv keys post testnet
			common.PanicIfError(config.SetOperatorKeys(config.GlobalCfg.OperatorKey))

			// create db Instance
			tempDB, err := db.NewDB()
			common.PanicIfError(err)
			// init global DB instance
			db.DBInstance = tempDB

			// create and init global config object
			types.ContractCallerObj, err = types.NewContractCaller()
			common.PanicIfError(err)

			//
			// Create all the required services
			//

			// create aggregator service
			aggregator := poller.NewAggregator(db.DBInstance)

			// create the syncer service
			syncer := listener.NewSyncer()

			//
			// init genesis state
			//

			// fetch number of batches submitted on-chain
			batchCount, err := types.ContractCallerObj.TotalBatches()
			common.PanicIfError(err)

			// if there are batches submitted already, start syncer
			// syncer will figure out if we are synced or not and will proceed accordingly
			if batchCount != 0 {
				// If !0, start syncer to sync all the blocks
				// TODO start syncer
			} else {
				// fetch 0 root from the contract
				root, err := types.ContractCallerObj.FetchBalanceTreeRoot()
				common.PanicIfError(err)

				// get the number of batches we have stored
				storedBatchCount, err := db.DBInstance.GetBatchCount()
				if err != nil {
					panic(err)
				}

				// figure out if we need to add adccounts
				if storedBatchCount == 0 {
					// persist batch info to DB
					err = db.DBInstance.InsertBatchInfo(root, uint64(batchCount))
					common.PanicIfError(err)
				}

				storedAccCount, err := db.DBInstance.GetAccountCount()
				if err != nil {
					panic(err)
				}

				// if there are no accounts add genesis accounts else skip
				if storedAccCount == 0 {
					// read genesis file and populate all accounts
					genAcc, err := config.ReadGenesisFile()
					common.PanicIfError(err)
					err = db.DBInstance.InsertGenAccounts(genAcc.Accounts)
					common.PanicIfError(err)
				}
			}

			// set last recorded block in DB for syncer
			llog, err := db.DBInstance.GetLastListenerLog()
			if err != nil && gorm.IsRecordNotFoundError(err) {
				// if no record is present so insert the record in config.toml
				err := db.DBInstance.StoreListenerLog(types.ListenerLog{LastRecordedBlock: cfg.LastRecordedBlock})
				common.PanicIfError(err)
			} else if err != nil && !gorm.IsRecordNotFoundError(err) {
				common.PanicIfError(err)
			}
			fmt.Println("Last recorded block for syncer present", llog.LastRecordedBlock)

			// go routine to catch signal
			catchSignal := make(chan os.Signal, 1)
			signal.Notify(catchSignal, os.Interrupt)
			go func() {
				// sig is a ^C, handle it
				for range catchSignal {
					aggregator.Stop()
					syncer.Stop()
					db.DBInstance.Close()

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

			if err := syncer.Start(); err != nil {
				log.Fatalln("Unable to start syncer", "error")
			}

			// TODO replace this with port from config
			http.ListenAndServe(":8080", r)
			fmt.Println("Server started on port 8080 🎉")
		},
	}
}
