package main

import (
	"crypto/ecdsa"
	"database/sql"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"os/signal"
	"strings"

	"github.com/BOPR/common"
	"github.com/BOPR/config"
	"github.com/BOPR/simulator"
	"github.com/ethereum/go-ethereum/crypto"
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
		Use:   "hubble",
		Short: "Optimistic Rollup Daemon (server)",
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
	rootCmd.AddCommand(StartSimulatorCmd())
	rootCmd.AddCommand(AddGenesisAcccountsCmd())
	rootCmd.AddCommand(SendTransferTx())
	rootCmd.AddCommand(CreateDatabase())
	rootCmd.AddCommand(CreateUsers())
	rootCmd.AddCommand(migrationCmd)

	executor := Executor{rootCmd, os.Exit}
	if err := executor.Command.Execute(); err != nil {
		fmt.Println("Error while executing command", err)
		return
	}
}

// ResetCmd resets all the collections
func ResetCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "reset",
		Short: "reset database",
		Run: func(cmd *cobra.Command, args []string) {

			err := config.ParseAndInitGlobalConfig("")
			common.PanicIfError(err)
			// TODO fix this command for mysql database
			// create new DB instance
			// dbInstance, err := db.NewDB()
			// defer dbInstance.Close()
			// common.PanicIfError(err)
			// fmt.Println("Resetting database", "db", common.DATABASE)
			// err = dbInstance.MgoSession.DropDatabase(common.DATABASE)
			// common.PanicIfError(err)
		},
	}
}

type UserList struct {
	Users []User `json:"users"`
}

type User struct {
	Address   string `json:"address"`
	PublicKey string `json:"pubkey"`
	PrivKey   string `json:"privkey"`
}

// CreateUsers creates the database
func CreateUsers() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "create-users",
		Short: "Create users to be used in simulations",
		RunE: func(cmd *cobra.Command, args []string) error {
			var users []User
			for i := 0; i < 3; i++ {
				privKey, err := crypto.GenerateKey()
				if err != nil {
					fmt.Println("Error generating private key", err)
					return err
				}
				publicKey := privKey.Public()
				ecsdaPubKey, ok := publicKey.(*ecdsa.PublicKey)
				if !ok {
					return errors.New("Unable to convert public key")
				}
				newUser := User{Address: crypto.PubkeyToAddress(*ecsdaPubKey).String(), PublicKey: "0x" + hex.EncodeToString(crypto.FromECDSAPub(ecsdaPubKey)), PrivKey: hex.EncodeToString(crypto.FromECDSA(privKey))}
				users = append(users, newUser)
			}
			bz, err := json.MarshalIndent(UserList{Users: users}, "", " ")
			if err != nil {
				return err
			}
			return ioutil.WriteFile("users.json", bz, 0644)
		},
	}
	// cmd.Flags().Int64(FlagNumberOfUsers, 2, "--count=<count>")
	// cmd.MarkFlagRequired(FlagNumberOfUsers)
	return cmd
}
func AddGenesisAcccountsCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "add-gen-accounts",
		Short: "Adds the accounts present in genesis account to the contract",
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
			// init global config
			config.GlobalCfg = cfg
		},
	}
}

// CreateDatabase creates the database
func CreateDatabase() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "create-database",
		Short: "Create a new database",
		RunE: func(cmd *cobra.Command, args []string) error {
			if err := config.ParseAndInitGlobalConfig(""); err != nil {
				return err
			}
			splitStrings := strings.Split(config.GlobalCfg.FormattedDBURL(), "/")
			connectionString := []string{splitStrings[0], "/"}
			dbNew, err := sql.Open("mysql", strings.Join(connectionString, ""))
			if err != nil {
				return err
			}
			defer dbNew.Close()
			_, err = dbNew.Exec(fmt.Sprintf("CREATE DATABASE IF NOT EXISTS %v", "hubble"))
			if err != nil {
				return err
			}
			return nil
		},
	}
	cmd.Flags().StringP(FlagDatabaseName, "", "", "--dbname=<database-name>")
	// cmd.MarkFlagRequired(FlagDatabaseName)
	return cmd
}

// StartSimulatorCmd starts the simulator
func StartSimulatorCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "start-simulating",
		Short: "starts a simulator that sends transaction to the rollupchain periodically",
		Run: func(cmd *cobra.Command, args []string) {
			sim := simulator.NewSimulator()
			if err := sim.Start(); err != nil {
				panic(err)
			}

			// go routine to catch signal
			catchSignal := make(chan os.Signal, 1)
			signal.Notify(catchSignal, os.Interrupt)
			go func() {
				for range catchSignal {
					sim.Stop()
					// exit
					os.Exit(1)
				}
			}()

			r := mux.NewRouter()
			err := http.ListenAndServe(":4000", r)
			if err != nil {
				panic(err)
			}
		},
	}
}
