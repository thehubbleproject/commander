package rest

import (
	"net/http"
)

type (
	// TxReceiver represents the tx received from user
	TxReceiver struct {
		To        string
		From      string
		Amount    uint64
		Nonce     uint64
		Fee       uint64
		TxType    uint64
		Signature string
	}
)

// TxReceiverHandler handles user txs
func TxReceiverHandler(w http.ResponseWriter, r *http.Request) {
	var tx TxReceiver
	if !ReadRESTReq(w, r, &tx) {
		WriteErrorResponse(w, http.StatusBadRequest, "Cannot read request")
	}
	return
}
