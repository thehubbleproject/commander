package rest

import (
	"net/http"
)

type (
	// TxReceiver represents the tx received from user
	TxReceiver struct {
		To        string `json:"to"`
		From      string `json:"from"`
		Amount    uint64
		Nonce     uint64
		fee       uint64
		Type      uint64
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
