package types

import (
	"encoding/hex"
	"math/big"

	ethCmn "github.com/ethereum/go-ethereum/common"
)

type ByteArray [32]byte

func HexToByteArray(a string) (b ByteArray, err error) {
	bz, err := hex.DecodeString(a)
	if err != nil {
		return b, err
	}
	return BytesToByteArray(bz), nil
}

func BytesToByteArray(bz []byte) ByteArray {
	var temp [32]byte
	copy(temp[:], bz)
	return temp
}

type Hash ethCmn.Hash
type Address ethCmn.Address

//
// utils
//
func ToUint64(i *big.Int) uint64 {
	return i.Uint64()
}

func UintToBigInt(a uint64) *big.Int {
	temp := big.NewInt(0)
	temp.SetUint64(a)
	return temp
}
