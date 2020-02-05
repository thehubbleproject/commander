package db

import (
	"fmt"
	"testing"

	"github.com/BOPR/db/mocks"
	"github.com/BOPR/types"
)

func TestFetchSiblings(t *testing.T) {
	dbMock := mocks.IDB{}
	// dbInstance, err := NewDB("")

	var paths []uint64
	paths = []uint64{11, 10, 01, 00}
	accounts := GenAccounts(paths)
	for i, path := range paths {
		dbMock.On("GetAccount", path).Return(accounts[i], nil)
	}
	fmt.Println("accounts created", accounts)
	fmt.Println(dbMock.GetAccount(paths[0]))
	data, err := FetchSiblings(paths[0], &dbMock)
	if err != nil {
		fmt.Println("error", err)
	}

	fmt.Printf("data is %v", data)
}

func GenAccounts(paths []uint64) (accounts []types.AccountLeaf) {
	for i := 0; i < 4; i++ {
		accounts = append(accounts, types.NewAccountLeaf(paths[i], 100, 1, 0))
	}
	return accounts
}

// func GenAccounts(num int) (accounts []types.AccountLeaf) {
// 	var height int
// 	if num%2 == 0 {
// 		height = num / 2
// 	} else {
// 		height = num/2 + 1
// 	}
// 	fmt.Println("height", height)
// 	var path []byte
// 	for i := 0; i < height; i++ {
// 		path = append(path, []byte("1")...)
// 	}
// 	fmt.Println("path generated 1", path)
// 	fmt.Println("path generated 2", string(path))
// 	pathInt, err := strconv.Atoi(string(path))
// 	if err != nil {
// 		fmt.Println("error", err)
// 	}

// 	pathStr := strconv.Itoa(pathInt)
// 	pathByte := []byte(pathStr)
// 	fmt.Println("path byte", pathByte)
// 	pathByte[0] = byte([]byte("0")...)
// 	pathInt, err = strconv.Atoi(string(pathByte))
// 	if err != nil {
// 		fmt.Println("error", err)
// 	}
// 	fmt.Println("post path byte", pathByte)

// 	fmt.Println("path generated", pathInt)

// 	for i := num; i > 0; i-- {
// 		accountPath := uint64(common.FlipBit(common.ExtractBit(pathInt, i)))
// 		fmt.Println("gen account path", accountPath)
// 		accounts = append(accounts, types.NewAccountLeaf(
// 			uint64(common.FlipBit(common.ExtractBit(pathInt, i))),
// 			100,
// 			1,
// 			0,
// 		))
// 	}
// 	return accounts
// }
