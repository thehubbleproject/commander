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

	"github.com/BOPR/poller"
	"github.com/BOPR/rest"
	"github.com/BOPR/types"
	"github.com/gorilla/mux"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

const (
	WithConfigPathFlag = "config-path"
	ConfigFileName     = "config"
)

// Executor wraps the cobra Command with a nicer Execute method
type Executor struct {
	*cobra.Command
	Exit func(int) // this is os.Exit by default, override in tests
}

func main() {
	rootCmd := &cobra.Command{
		Use:   "bopr",
		Short: "Beacon Optimistic Rollup Daemon (server)",
	}

	// add new persistent flag for heimdall-config
	rootCmd.PersistentFlags().String(
		WithConfigPathFlag,
		"",
		"Config file path (default ./config.toml)",
	)

	// bind with-heimdall-config config with root cmd
	viper.BindPFlag(
		WithConfigPathFlag,
		rootCmd.Flags().Lookup(WithConfigPathFlag),
	)
	rootCmd.AddCommand(InitCmd())
	rootCmd.AddCommand(StartCmd())
	rootCmd.AddCommand(ResetCmd())

	executor := Executor{rootCmd, os.Exit}
	if err := executor.Command.Execute(); err != nil {
		fmt.Println("Error while executing command", err)
		return
	}
}

// InitCmd generated init command to initialise the config file
func InitCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "init",
		Short: "Initialises Configration for BOPR",
		Run: func(cmd *cobra.Command, args []string) {
			defaultConfig := config.GetDefaultConfig()
			config.WriteConfigFile("./config.toml", &defaultConfig)
			gen := config.DefaultGenesisAccounts()
			if err := config.WriteGenesisFile(gen); err != nil {
				panic(err)
			}
		},
	}
}

// ResetCmd resets all the collections
func ResetCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "reset",
		Short: "reset database",
		Run: func(cmd *cobra.Command, args []string) {
			viperObj := viper.New()
			dir, err := os.Getwd()
			common.PanicIfError(err)

			viperObj.SetConfigName(ConfigFileName) // name of config file (without extension)
			viperObj.AddConfigPath(dir)
			err = viperObj.ReadInConfig()
			common.PanicIfError(err)

			var cfg config.Configuration
			if err = viperObj.UnmarshalExact(&cfg); err != nil {
				common.PanicIfError(err)
			}
			config.GlobalCfg = cfg

			// create mgo session
			session, err := db.NewSession(cfg.MongoDB)
			common.PanicIfError(err)
			fmt.Println("Resetting database", "db", common.DATABASE)
			err = session.DropDatabase(common.DATABASE)
			common.PanicIfError(err)
		},
	}
}

// StartCmd starts the daemon
func StartCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "start",
		Short: "Starts BOPR daemon",
		Run: func(cmd *cobra.Command, args []string) {
			viperObj := viper.New()
			dir, err := os.Getwd()
			common.PanicIfError(err)

			viperObj.SetConfigName(ConfigFileName) // name of config file (without extension)
			viperObj.AddConfigPath(dir)
			err = viperObj.ReadInConfig()
			common.PanicIfError(err)

			var cfg config.Configuration
			if err = viperObj.UnmarshalExact(&cfg); err != nil {
				common.PanicIfError(err)
			}

			config.GlobalCfg = cfg

			// create mgo session
			session, err := db.NewSession(cfg.MongoDB)
			common.PanicIfError(err)

			db.MgoSession = *session
			aggregator := poller.NewAggregator()
			types.ContractCallerObj, err = types.NewContractCaller()
			common.PanicIfError(err)

			// fetch number of batches
			batchCount, err := types.ContractCallerObj.TotalBatches()
			common.PanicIfError(err)

			// If 0, dump to DB and start building on itd
			if batchCount == 0 {
				root, err := types.ContractCallerObj.FetchBalanceTreeRoot()
				common.PanicIfError(err)

				// persist batch info to DB
				err = db.InsertBatchInfo(root, batchCount)
				common.PanicIfError(err)
			}

			// If !0, start syncer to sync all the blocks
			// TODO start syncer

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

			// TODO replace this with port from config
			http.ListenAndServe(":8080", r)
			fmt.Println("Server started on port 8080 🎉")
		},
	}
}
