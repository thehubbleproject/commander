package main

import (
	"encoding/hex"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"golang.org/x/crypto/sha3"
	"log"
	"math/big"
)

func main() {
	uint256Ty, _ := abi.NewType("uint256", "uint256", nil)
	// bytes32Ty, _ := abi.NewType("bytes32", "bytes32", nil)
	addressTy, _ := abi.NewType("address", "address", nil)

	arguments := abi.Arguments{
		{
			Type: addressTy,
		},
		{
			Type: addressTy,
		},
		{
			Type: uint256Ty,
		},
	}

	bytes, _ := arguments.Pack(
		common.HexToAddress("0x0000000000000000000000000000000000000001"),
		common.HexToAddress("0x0000000000000000000000000000000000000001"),
		big.NewInt(42),
	)
	var buf []byte
	log.Println("dataBytes", hex.EncodeToString(bytes))
	hash := sha3.NewLegacyKeccak256()
	hash.Write(bytes)
	buf = hash.Sum(buf)

	log.Println(hexutil.Encode(buf))
	// output:
	// 0x1f214438d7c061ad56f98540db9a082d372df1ba9a3c96367f0103aa16c2fe9a
}
