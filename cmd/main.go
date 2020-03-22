package main

import (
	"fmt"
	"os"

	"github.com/BOPR/common"
	"github.com/BOPR/config"

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
	rootCmd.AddCommand(AddGenesisAcccountsCmd())
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

			err := config.ParseAndInitGlobalConfig()
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
			// genAccs, err := config.ReadGenesisFile()
			// common.PanicIfError(err)
			// contractCaller, err := types.NewContractCaller()
			// for _, genAcc := range genAccs.Accounts {
			// 	err := contractCaller.AddAccount(types.NewUserAccount(genAcc.Path, genAcc.Balance, genAcc.TokenType, genAcc.Nonce))
			// 	common.PanicIfError(err)
			// }
			// common.PanicIfError(err)
		},
	}
}
