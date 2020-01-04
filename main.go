package main

import (
	"encoding/hex"
	"fmt"

	"github.com/BOPR/types"
	"github.com/cbergoon/merkletree"
)

func main() {
	// create 2 leaves

	// create merkel root

	// match with the one in solidity

	// add the testcase to test-suite

	// leaf1 := types.NewAccountLeaf(0, 10)
	// leaf2 := types.NewAccountLeaf(1, 10)

	var list []merkletree.Content
	src := []byte("0x2aa012a32db4297b6c1ec06b81e498154b4e8d46")
	encodedStr := hex.EncodeToString(src)

	dataBytes, err := hex.DecodeString(encodedStr)
	if err != nil {
		fmt.Println("data bytes", dataBytes, "error", err)
	}

	data := types.LeafData(dataBytes)
	list = append(list, data)
	// list = append(list, data)
	fmt.Println("Input dataBytes", list, " Expected ", dataBytes)
	tree, err := types.CreateTree(list)
	if err != nil {
		fmt.Printf("error", err)
	}
	fmt.Println("root", hex.EncodeToString(tree.MerkleRoot()))
}
