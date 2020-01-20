package main

import (
	"encoding/hex"
	"fmt"

	crypto "github.com/ethereum/go-ethereum/crypto"
)

func main() {
	bz, err := hex.DecodeString(hex.EncodeToString([]byte("0x12")))
	if err != nil {
		panic(err)
	}
	hash := crypto.Keccak256(bz)

	fmt.Println("hash", hex.EncodeToString(hash))
}
