package types

import (
	"github.com/BOPR/contracts/rollup"
)

type AccountLeaf struct {
	Path      uint64
	Balance   uint64
	TokenType uint64
	Nonce     uint64
}

func NewAccountLeaf(path, balance, tokenType, nonce uint64) AccountLeaf {
	return AccountLeaf{
		Path:      path,
		Balance:   balance,
		TokenType: tokenType,
		Nonce:     nonce,
	}
}

func (acc *AccountLeaf) ToABIAccount() rollup.DataTypesAccount {
	return rollup.DataTypesAccount{
		Path:      UintToBigInt(acc.Path),
		Balance:   UintToBigInt(acc.Balance),
		TokenType: UintToBigInt(acc.TokenType),
		Nonce:     UintToBigInt(acc.Nonce),
	}
}
