package main

import (
	"fmt"

	"github.com/BOPR/core"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

//  SendTransferTx generated init command to initialise the config file
func SendTransferTx() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "transfer",
		Short: "Transfers assets between 2 accounts",
		RunE: func(cmd *cobra.Command, args []string) error {
			toIndex := viper.GetUint64(FlagToAccountID)
			fromIndex := viper.GetUint64(FlagFromAccountID)

			tokenID := viper.GetUint64(FlagTokenID)
			nonce := viper.GetUint64(FlagNonce)

			sig := viper.GetString(FlagSignature)

			amount := viper.GetUint64(FlagAmount)
			tx := core.NewPendingTx(toIndex, fromIndex, amount, nonce, sig, tokenID)
			tx.AssignHash()

			db, err := core.NewDB()
			if err != nil {
				return err
			}
			defer db.Close()
			err = db.InsertTx(&tx)
			if err != nil {
				return err
			}
			fmt.Println("Transaction submitted successfully", "hash", tx.TxHash)
			return nil
		},
	}
	cmd.Flags().StringP(FlagToAccountID, "", "", "--to=<to-account>")
	cmd.Flags().StringP(FlagFromAccountID, "", "", "--from=<from-account>")

	cmd.Flags().StringP(FlagTokenID, "", "", "--token=<token-id>")

	cmd.Flags().StringP(FlagNonce, "", "", "--nonce=<nonce>")
	cmd.Flags().StringP(FlagSignature, "", "", "--sig=<signature>")

	cmd.Flags().StringP(FlagAmount, "", "", "--amount=<amount>")
	cmd.MarkFlagRequired(FlagNonce)

	cmd.MarkFlagRequired(FlagTokenID)
	return cmd
}

//  SendTransferTx generated init command to initialise the config file
func SendDummyTransfers() *cobra.Command {
	return &cobra.Command{
		Use:   "dummy-transfer",
		Short: "Transfers assets between 2 accounts",
		Run: func(cmd *cobra.Command, args []string) {

		},
	}
}
