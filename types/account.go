package types

import (
	"github.com/BOPR/contracts/rollup"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"math/big"
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

func (acc *AccountLeaf) ABIEncode() []byte {
	uint256Ty, _ := abi.NewType("uint256", "uint256", nil)
	arguments := abi.Arguments{
		{
			Type: uint256Ty,
		},
		{
			Type: uint256Ty,
		},
		{
			Type: uint256Ty,
		},
		{
			Type: uint256Ty,
		},
	}
	bytes, _ := arguments.Pack(
		big.NewInt(int64(acc.Path)),
		big.NewInt(int64(acc.Balance)),
		big.NewInt(int64(acc.TokenType)),
		big.NewInt(int64(acc.Nonce)),
	)
	return bytes
}
