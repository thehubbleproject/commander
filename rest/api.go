package rest

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/BOPR/core"
)

type (
	// TxReceiver represents the tx received from user
	TxReceiver struct {
		From      uint64 `json:"from"`
		To        uint64 `json:"to"`
		Message   []byte `json:"message"`
		Signature string `json:"sig"`
	}

	TransferTx struct {
		From      uint64 `json:"from"`
		To        uint64 `json:"to"`
		Amount    uint64 `json:"amount"`
		Nonce     uint64 `json:"nonce"`
		TokenType uint64 `json:"token"`
		TxType    uint64 `json:"txType"`
	}
)

// TxReceiverHandler handles user txs
func TxReceiverHandler(w http.ResponseWriter, r *http.Request) {
	// receive the payload and read
	var tx TxReceiver
	if !ReadRESTReq(w, r, &tx) {
		WriteErrorResponse(w, http.StatusBadRequest, "Cannot read request")
	}

	// create a new pending transaction
	userTx := core.NewPendingTx(tx.From, tx.To, core.TX_TRANSFER_TYPE, tx.Signature, tx.Message)

	// add the transaction to pool
	err := core.DBInstance.InsertTx(&userTx)
	if err != nil {
		WriteErrorResponse(w, http.StatusBadRequest, "Cannot read request")
	}

	output, err := json.Marshal(userTx)
	if err != nil {
		WriteErrorResponse(w, http.StatusBadRequest, "Unable to marshall account")
	}

	// write headers and data
	w.Header().Set("Content-Type", "application/json")
	_, _ = w.Write(output)
	return
}

// GetAccountHandler fetches the user account data like balance, token type and nonce
func GetAccountHandler(w http.ResponseWriter, r *http.Request) {
	vars := r.URL.Query()
	IDstr := vars.Get("ID")
	ID, err := strconv.ParseUint(IDstr, 0, 64)
	if err != nil {
		WriteErrorResponse(w, http.StatusBadRequest, "Invalid ID")
	}
	fmt.Println(ID)
	var account core.UserAccount
	// account, err := core.DBInstance.GetAccount(ID)
	// if err != nil {
	// 	WriteErrorResponse(w, http.StatusBadRequest, fmt.Sprintf("Account with ID %v not found", ID))
	// }
	output, err := json.Marshal(account)
	if err != nil {
		WriteErrorResponse(w, http.StatusBadRequest, "Unable to marshall account")
	}
	w.Header().Set("Content-Type", "application/json")
	_, _ = w.Write(output)
	return
}

// GetAccountHandler fetches the user account data like balance, token type and nonce
func GetTxTemplateForTxType(w http.ResponseWriter, r *http.Request) {
	// vars := r.URL.Query()
	// typeStr := vars.Get("txType")
	// txType, err := strconv.ParseUint(typeStr, 0, 64)
	// if err != nil {
	// 	WriteErrorResponse(w, http.StatusBadRequest, "Invalid txType")
	// }
	var account core.UserAccount
	// account, err := core.DBInstance.GetAccount(ID)
	// if err != nil {
	// 	WriteErrorResponse(w, http.StatusBadRequest, fmt.Sprintf("Account with ID %v not found", ID))
	// }
	output, err := json.Marshal(account)
	if err != nil {
		WriteErrorResponse(w, http.StatusBadRequest, "Unable to marshall account")
	}
	w.Header().Set("Content-Type", "application/json")
	_, _ = w.Write(output)
	return
}
