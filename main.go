package main

import (
	"context"
	"encoding/hex"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
	ethCmn "github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/rpc"

	"github.com/BOPR/config"
	"github.com/BOPR/listener"
	"github.com/BOPR/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/gorilla/mux"
)

func main() {
	TestCallData()
}

func TestCallData() {
	RPCClient, err := rpc.Dial("https://goerli.infura.io/v3/7e99d705c25844b59df18449632dd97c")
	client := ethclient.NewClient(RPCClient)
	tx, _, err := client.TransactionByHash(context.Background(), ethCmn.HexToHash("0x92571ece7656376802d5e49a67b80ab51a9bfd15a451564b9e29e13928338f7f"))
	if err != nil {
		panic(err)
	}
	fmt.Println("transaction", hex.EncodeToString(tx.Data()))

	myAbi, err := abi.JSON(strings.NewReader("[{\"constant\":false,\"inputs\":[{\"internalType\":\"bytes[]\",\"name\":\"_txs\",\"type\":\"bytes[]\"}],\"name\":\"Batch\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"))
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("length", len(tx.Data()))
	payload := tx.Data()
	decodedPayload := payload[4:]
	inputDataMsap := make(map[string]interface{})
	method := myAbi.Methods["Batch"]
	err = method.Inputs.UnpackIntoMap(inputDataMsap, decodedPayload)
	if err != nil {
		log.Fatal("error decoding", err)
	}
	fmt.Println("decoded data", hex.EncodeToString(inputDataMsap["_txs"].([][]uint8)[0]))
}

func TestSyncer() {
	config.ParseAndInitGlobalConfig()
	syncer := listener.NewSyncer()
	syncer.DBInstance.StoreListenerLog(types.ListenerLog{LastRecordedBlock: "101"})
	if err := syncer.Start(); err != nil {
		log.Fatalln("Unable to start syncer", "error")
	}

	// go routine to catch signal
	catchSignal := make(chan os.Signal, 1)
	signal.Notify(catchSignal, os.Interrupt)
	go func() {
		// sig is a ^C, handle it
		for range catchSignal {
			syncer.Stop()
			// exit
			os.Exit(1)
		}
	}()
	r := mux.NewRouter()
	http.ListenAndServe(":8080", r)
}
