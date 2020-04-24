package main

import (
	"fmt"

	"github.com/BOPR/core"
)

func main() {
	TestUserAccountUpdate()
}

func TestTrim() {
	fmt.Println(core.TrimPathToParentPath("1001"))
}
func TestUserAccountUpdate() {
	db, err := core.NewDB()
	if err != nil {
		fmt.Println("error creating database")
		panic(err)

	}
	fmt.Println("created DB", db)
	var newAcc core.UserAccount
	// originalAccountPath := "100"
	newAcc.Status = 100
	newAcc.AccountID = 100
	fmt.Println("new accout", newAcc)
	err = db.Instance.Create(&newAcc).Error
	if err != nil {
		panic(err)
	}
	var user core.UserAccount
	err = db.Instance.Where("path = ?", "").First(&user).Error
	fmt.Println("user", user, err)

	// newAcc.AccountID = 1001
	// newAcc.Path = "10101"
	// err = db.Instance.Create(&newAcc).Error
	// if err != nil {
	// 	panic(err)
	// }
	// newAcc.AccountID = 102
	// newAcc.Path = "01"
	// err = db.Instance.Model(&newAcc).Where("path = ?", originalAccountPath).Update(&newAcc).Error
	// if err != nil {
	// 	panic(err)
	// }
}

// func TestABIEncodeAndDecode() {
// 	RPCClient, err := rpc.Dial("")
// 	client := ethclient.NewClient(RPCClient)
// 	rollupContractAddress := ethCmn.HexToAddress("0x61b14bAA77069494fCff00EAeeCf0212d5Ac1d10")
// 	instance, err := trial.NewTrial(rollupContractAddress, client)
// 	if err != nil {
// 		panic(err)
// 	}

// 	rollupABI, err := abi.JSON(strings.NewReader(trial.TrialABI))
// 	from := trial.TrialUserAccount{
// 		ID:        big.NewInt(100),
// 		Balance:   big.NewInt(100),
// 		TokenType: big.NewInt(100),
// 		Nonce:     big.NewInt(100),
// 	}

// 	txData := trial.TrialTransaction{
// 		From:      from,
// 		To:        from,
// 		TokenType: big.NewInt(100),
// 		Amount:    100,
// 		Signature: []byte("dsds"),
// 	}

// 	data, err := rollupABI.Pack("ABIEncodeTransaction", txData)
// 	if err != nil {
// 		fmt.Println("Unable to pack tx for approve", "error", err)
// 	}

// 	// generate call msg
// 	callMsg := ethereum.CallMsg{
// 		To:   &rollupContractAddress,
// 		Data: data,
// 	}

// 	auth, err := bazooka.GenerateAuthObj(client, callMsg)
// 	fmt.Println("auth generated", auth, err)

// 	tx, err := instance.ABIEncodeTransaction(auth, txData)
// 	if err != nil {
// 		fmt.Println("Unable to send tx", err)
// 	}
// 	fmt.Println("sent transaction", tx.Hash().String())
// }

// func TestCallData() {
// 	RPCClient, err := rpc.Dial("")
// 	client := ethclient.NewClient(RPCClient)
// 	tx, _, err := client.TransactionByHash(context.Background(), ethCmn.HexToHash("0x16325b88c533d8e75be9efefa2e96a39c75e106ec6d0974da1f143702b181809"))
// 	if err != nil {
// 		panic(err)
// 	}
// 	fmt.Println("transaction", hex.EncodeToString(tx.Data()))

// 	myAbi, err := abi.JSON(strings.NewReader("[{\"constant\":false,\"inputs\":[{\"internalType\":\"bytes[]\",\"name\":\"_txs\",\"type\":\"bytes[]\"}],\"name\":\"Batch\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"))
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	fmt.Println("length", len(tx.Data()))
// 	payload := tx.Data()
// 	decodedPayload := payload[4:]
// 	inputDataMsap := make(map[string]interface{})
// 	method := myAbi.Methods["Batch"]
// 	err = method.Inputs.UnpackIntoMap(inputDataMsap, decodedPayload)
// 	if err != nil {
// 		log.Fatal("error decoding", err)
// 	}

// 	fmt.Println("decoded data", hex.EncodeToString(inputDataMsap["_txs"].([][]byte)[2]))
// }

// func TestSyncer() {
// 	config.ParseAndInitGlobalConfig()
// 	syncer := listener.NewSyncer()
// 	batch := types.Batch{StakeAmount: 100}
// 	fmt.Println(syncer.DBInstance.AddNewBatch(batch))
// 	syncer.DBInstance.StoreListenerLog(types.ListenerLog{LastRecordedBlock: "101"})
// 	if err := syncer.Start(); err != nil {
// 		log.Fatalln("Unable to start syncer", "error")
// 	}

// 	// go routine to catch signal
// 	catchSignal := make(chan os.Signal, 1)
// 	signal.Notify(catchSignal, os.Interrupt)
// 	go func() {
// 		// sig is a ^C, handle it
// 		for range catchSignal {
// 			syncer.Stop()
// 			// exit
// 			os.Exit(1)
// 		}
// 	}()
// 	r := mux.NewRouter()
// 	http.ListenAndServe(":8080", r)
// }
