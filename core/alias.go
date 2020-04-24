package core

import (
	"encoding/hex"

	ethCmn "github.com/ethereum/go-ethereum/common"
)

type Hash ethCmn.Hash
type Address ethCmn.Address

type ByteArray [32]byte

func (b ByteArray) String() string {
	return hex.EncodeToString(b[:])
}

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
