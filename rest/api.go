package rest

import (
	"fmt"
	"net/http"

	"github.com/BOPR/types"
)

type (
	// TxReceiver represents the tx received from user
	TxReceiver struct {
		To        uint64 `json:"to"`
		From      uint64 `json:"from"`
		Amount    uint64 `json:"amount"`
		Nonce     uint64 `json:"nonce"`
		Fee       uint64 `json:"fee"`
		TxType    uint64 `json:"type"`
		Signature string `json:"sig"`
	}
)

// TxReceiverHandler handles user txs
func TxReceiverHandler(w http.ResponseWriter, r *http.Request) {
	var tx TxReceiver
	if !ReadRESTReq(w, r, &tx) {
		WriteErrorResponse(w, http.StatusBadRequest, "Cannot read request")
	}
	fmt.Printf("transaction", tx)
	userTx := types.NewTx(tx.To, tx.From, tx.Amount, tx.Nonce, tx.Signature)
	userTx.Insert()
	return
}
