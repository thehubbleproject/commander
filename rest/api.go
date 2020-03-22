package rest

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	db "github.com/BOPR/db"
	"github.com/BOPR/types"
)

type (
	// TxReceiver represents the tx received from user
	TxReceiver struct {
		To        uint64 `json:"to"`
		From      uint64 `json:"from"`
		Amount    uint64 `json:"amount"`
		Nonce     uint64 `json:"nonce"`
		TokenID   uint64 `json:"token"`
		Signature string `json:"sig"`
	}
)

// GetAccountHandler fetches the user account data like balance, token type and nonce
func GetAccountHandler(w http.ResponseWriter, r *http.Request) {
	vars := r.URL.Query()
	IDstr := vars.Get("ID")
	ID, err := strconv.ParseUint(IDstr, 0, 64)
	if err != nil {
		WriteErrorResponse(w, http.StatusBadRequest, "Invalid ID")
	}

	account, err := db.DBInstance.GetAccount(ID)
	if err != nil {
		WriteErrorResponse(w, http.StatusBadRequest, fmt.Sprintf("Account with ID %v not found", ID))
	}
	output, err := json.Marshal(account)
	if err != nil {
		WriteErrorResponse(w, http.StatusBadRequest, "Unable to marshall account")
	}
	w.Header().Set("Content-Type", "application/json")
	_, _ = w.Write(output)
	return
}

// TxReceiverHandler handles user txs
func TxReceiverHandler(w http.ResponseWriter, r *http.Request) {
	// receive the payload and read
	var tx TxReceiver
	if !ReadRESTReq(w, r, &tx) {
		WriteErrorResponse(w, http.StatusBadRequest, "Cannot read request")
	}
	// create a new pending transaction
	userTx := types.NewPendingTx(tx.To, tx.From, tx.Amount, tx.Nonce, tx.Signature, tx.TokenID)

	// do basic input validations
	err := userTx.ValidateBasic()
	if err != nil {
		WriteErrorResponse(w, http.StatusBadRequest, "Cannot read request")
	}

	// assign the transaction a HASH
	userTx.AssignHash()

	// add the transaction to pool
	db.DBInstance.InsertTx(&userTx)

	output, err := json.Marshal(userTx)
	if err != nil {
		WriteErrorResponse(w, http.StatusBadRequest, "Unable to marshall account")
	}

	// write headers and data
	w.Header().Set("Content-Type", "application/json")
	_, _ = w.Write(output)

	return
}
