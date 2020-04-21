package types

import (
	"bytes"
	"fmt"

	"github.com/cbergoon/merkletree"
	merkle "github.com/cbergoon/merkletree"
	"golang.org/x/crypto/sha3"
)

type DepositTree struct {
	Height           uint64
	NumberOfDeposits uint64
	Root             string
}

func (t *DepositTree) String() string {
	return fmt.Sprintf("DepositTree: H:%v Count:%v Root:%v", t.Height, t.NumberOfDeposits, t.Root)
}

type LeafData []byte

//CalculateHash hashes the values of a TestContent
func (data LeafData) CalculateHash() ([]byte, error) {
	h := sha3.NewLegacyKeccak256()
	// h := sha256.New()
	// if _, err := h.Write([]byte(data)); err != nil {
	// 	return nil, err
	// }
	return h.Sum(data), nil
}

//Equals tests for equality of two Contents
func (data LeafData) Equals(other merkletree.Content) (bool, error) {
	return bytes.Equal(data, other.(LeafData)), nil
}

// CreateTree creates a tree from a given list of contents
func CreateTree(list []merkle.Content) (tree *merkle.MerkleTree, err error) {
	t, err := merkle.NewTreeWithHashStrategy(list, sha3.NewLegacyKeccak256)
	if err != nil {
		return tree, err
	}
	return t, nil
}
