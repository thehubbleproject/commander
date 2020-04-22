package types

import (
	"encoding/hex"
	"math/big"
	"strconv"

	ethCmn "github.com/ethereum/go-ethereum/common"
)

type ByteArray [32]byte

type Hash ethCmn.Hash
type Address ethCmn.Address

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

//
// utils
//
func ToUint64(i *big.Int) uint64 {
	return i.Uint64()
}

func StringToBigInt(s string) *big.Int {
	t := big.NewInt(0)
	t.SetString(s, 2)
	return t
}

func UintToBigInt(a uint64) *big.Int {
	t := big.NewInt(0)
	t.SetUint64(a)
	return t
}

func GetBit(a uint64, index int) uint {
	b := UintToBigInt(a)
	return b.Bit(index)
}

func StringToUint(s string) (uint64, error) {
	i, err := strconv.ParseInt(s, 2, 64)
	if err != nil {
		return 0, err
	}
	return uint64(i), nil
}

func FlipBitInString(s string, i int) string {
	dataRune := []rune(s)
	if dataRune[i] == 48 {
		dataRune[i] = 49
	} else {
		dataRune[i] = 48
	}
	return string(dataRune)
}

// // GetChildrenPaths returns all the children leaf nodes
// func GetChildrenPaths(path string, maxDepth int) []string {

// }
