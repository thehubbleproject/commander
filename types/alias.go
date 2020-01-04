package types

import (
	ethCmn "github.com/ethereum/go-ethereum/common"
	"math/big"
)

type ByteArray [32]byte
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
