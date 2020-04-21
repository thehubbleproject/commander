package main

import (
	"fmt"
	"log"
	"math"
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
	"github.com/BOPR/types/bazooka"
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// StartCmd starts the daemon
func StartCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "start",
		Short: "Starts hubble daemon",
		Run: func(cmd *cobra.Command, args []string) {
			var err error
			// populate global config object
			ReadAndInitGlobalConfig()

			InitGlobalDBInstance()

			InitGlobalBazooka()

			InitDepositTree()

			logger := common.Logger.With("module", "hubble")

			//
			// Create all the required services
			//

			// create aggregator service
			aggregator := poller.NewAggregator(db.DBInstance)

			// create the syncer service
			syncer := listener.NewSyncer()

			// if no row is found then we are starting the node for the first time
			syncStatus, err := db.DBInstance.GetSyncStatus()
			if err != nil && gorm.IsRecordNotFoundError(err) {
				// read genesis file
				genesis, err := config.ReadGenesisFile()
				common.PanicIfError(err)

				// loads genesis data to the database
				LoadGenesisData(genesis)
			} else if err != nil && !gorm.IsRecordNotFoundError(err) {
				logger.Error("Error connecting to database", "error", err)
				common.PanicIfError(err)
			}

			logger.Info("Starting coordinator with sync and aggregator enabled", "lastSyncedEthBlock",
				syncStatus.LastEthBlockBigInt().String(),
				"lastSyncedBatch", syncStatus.LastBatchRecorded)

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
			r.HandleFunc("/account", rest.GetAccountHandler).Methods("GET")
			http.Handle("/", r)

			if err := syncer.Start(); err != nil {
				log.Fatalln("Unable to start syncer", "error")
			}

			if err := aggregator.Start(); err != nil {
				log.Fatalln("Unable to start aggregator", "error", err)
			}
			// TODO replace this with port from config
			http.ListenAndServe(":8080", r)
			fmt.Println("Server started on port 8080 🎉")
		},
	}
}

func ReadAndInitGlobalConfig() {
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

	// init global config
	config.GlobalCfg = cfg
	// TODO use a better way to handle priv keys post testnet
	common.PanicIfError(config.SetOperatorKeys(config.GlobalCfg.OperatorKey))
}

func InitGlobalDBInstance() {
	// create db Instance
	tempDB, err := db.NewDB()
	common.PanicIfError(err)

	// init global DB instance
	db.DBInstance = tempDB
}

func InitGlobalBazooka() {
	var err error
	// create and init global config object
	bazooka.LoadedBazooka, err = bazooka.NewPreLoadedBazooka()
	common.PanicIfError(err)
}

// LoadGenesisData helps load the genesis data into the DB
func LoadGenesisData(genesis config.Genesis) {
	diff := int(math.Exp2(float64(genesis.MaxTreeDepth))) - len(genesis.GenesisAccounts.Accounts)
	if diff < 0 {
		err := fmt.Errorf("Tree depth not enough to accomodate all geneiss account. Depth: %v NumberOfAccounts: %v", genesis.MaxTreeDepth, len(genesis.GenesisAccounts.Accounts))
		common.PanicIfError(err)
	}
	// fill the tree with zero leaves
	for diff > 0 {
		lastGenAcc := genesis.GenesisAccounts.Accounts[len(genesis.GenesisAccounts.Accounts)-1]
		newAcc := config.EmptyGenesisAccount()
		newAcc.Path = lastGenAcc.Path + 1
		genesis.GenesisAccounts.Accounts = append(genesis.GenesisAccounts.Accounts, newAcc)
		diff--
	}

	// load accounts
	err := db.DBInstance.InsertGenAccounts(genesis.GenesisAccounts.Accounts)
	common.PanicIfError(err)

	// load params
	newParams := types.Params{StakeAmount: genesis.StakeAmount, MaxDepth: genesis.MaxTreeDepth, MaxDepositSubTreeHeight: genesis.MaxDepositSubTreeHeight}
	db.DBInstance.UpdateStakeAmount(newParams.StakeAmount)
	db.DBInstance.UpdateMaxDepth(newParams.MaxDepth)
	db.DBInstance.UpdateDepositSubTreeHeight(newParams.MaxDepositSubTreeHeight)

	// load sync status
	db.DBInstance.UpdateSyncStatusWithBlockNumber(genesis.StartEthBlock)
	db.DBInstance.UpdateSyncStatusWithBatchNumber(0)
}

func InitDepositTree() {
	err := db.DBInstance.InitEmptyDepositTree()
	if err != nil {
		common.PanicIfError(err)
	}
}
