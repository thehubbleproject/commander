package types

import (
	"bytes"
	"crypto/sha256"

	"github.com/cbergoon/merkletree"
	merkle "github.com/cbergoon/merkletree"
)

type LeafData []byte

//CalculateHash hashes the values of a TestContent
func (data LeafData) CalculateHash() ([]byte, error) {
	h := sha256.New()
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
	t, err := merkle.NewTree(list)
	if err != nil {
		return tree, err
	}
	return t, nil
}
