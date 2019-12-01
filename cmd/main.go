package main

import (
	"fmt"
	"os"

	"github.com/BOPR/config"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

const (
	WithConfigPathFlag = "config-path"
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
	executor := Executor{rootCmd, os.Exit}
	if err := executor.Command.Execute(); err != nil {
		fmt.Println("Error while executing command", err)
		return
	}
	fmt.Println("Done!")
}

// InitCmd generated init command to initialise the config file
func InitCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "init",
		Short: "Initialises Configration for BOPR",
		Run: func(cmd *cobra.Command, args []string) {
			defaultConfig := config.GetDefaultConfig()
			config.WriteConfigFile("./config.toml", &defaultConfig)
		},
	}
}
