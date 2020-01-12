package main

import (
	"encoding/hex"
	"fmt"
	"math/big"
	"strconv"
	"strings"
)

func main() {
	// parentInt := StringToUint("0x11")
	// leftInt := StringToUint("0x0111111111111111111111111111111111111111111111111111111111111111")
	// fmt.Println("after left", parentInt&leftInt)
	// rightInt := StringToUint("0x1000000000000000000000000000000000000000000000000000000000000000")

	// var bigNum1 big.Int
	// var bigNum2 big.Int
	// var result big.Int

	// if _, ok := bigNum1.SetString("11", 2); !ok {
	// 	panic("invalid num1")
	// }
	// if _, ok := bigNum2.SetString("1000000000000000000000000000000000000000000000000000000000000000", 2); !ok {
	// 	panic("invalid num2")
	// }
	// result.Or(&bigNum1, &bigNum2)
	// fmt.Println("=>", result)
	// for i := result.BitLen() - 1; i >= 0; i-- {
	// 	fmt.Print(result.Bit(i))
	// }
	// resultsInt := parentInt | rightInt
	// fmt.Println("left int", resultsInt)
	// var bz = make([]byte, 100)

	// bitutil.ORBytes(bz, DecodeString("11"), DecodeString("1000000000000000000000000000000000000000000000000000000000000000"))
	// fmt.Println(hex.EncodeToString(bz))
	ShiftRight(uint64(101), 2)
}

func StringToUint(hexStr string) uint64 {
	// remove 0x suffix if found in the input string
	cleaned := strings.Replace(hexStr, "0x", "", -1)
	// base 16 for hexadecimal
	result, _ := strconv.ParseUint(cleaned, 16, 64)
	return uint64(result)
}

func DecodeString(hexString string) []byte {
	bz, err := hex.DecodeString(hexString)
	if err != nil {
		fmt.Println(err)
	}
	return bz
}

func ShiftRight(data uint64, times uint) {
	var bigNum1 big.Int
	bigNum1.SetUint64(data)
	result := bigNum1.Lsh(&bigNum1, times)
	fmt.Println("result is", result.String())
}
