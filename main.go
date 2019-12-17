package main

import (
	"encoding/hex"
	"fmt"

	"github.com/BOPR/common"
	"github.com/BOPR/types"
)

func main() {
	str := "Hello from ADMFactory.com"
	hx := hex.EncodeToString([]byte(str))
	tx := types.NewTx(1, 2, 3, 4, string(hx))
	minTx, err := tx.MinimalTx()
	common.PanicIfError(err)

	data := minTx.Serialise()
	fmt.Println("data", hex.EncodeToString(data))
	fmt.Println(types.DeserialiseMinimalTx(data))
}
