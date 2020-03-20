package types

import (
	"math/big"

	"github.com/BOPR/common"
	"github.com/BOPR/contracts/rollup"
	"github.com/ethereum/go-ethereum/accounts/abi"
)

type UserAccount struct {
	Path      uint64
	Balance   uint64
	TokenType uint64
	Nonce     uint64
}

func NewUserAccount(path, balance, tokenType, nonce uint64) UserAccount {
	return UserAccount{
		Path:      path,
		Balance:   balance,
		TokenType: tokenType,
		Nonce:     nonce,
	}
}

func (acc *UserAccount) ToABIAccount() rollup.DataTypesAccount {
	return rollup.DataTypesAccount{
		Path:      UintToBigInt(acc.Path),
		Balance:   UintToBigInt(acc.Balance),
		TokenType: UintToBigInt(acc.TokenType),
		Nonce:     UintToBigInt(acc.Nonce),
	}
}

func (acc *UserAccount) ABIEncode() []byte {
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

func AccsToLeafHashes(accs []UserAccount) (result [][32]byte) {
	for i, acc := range accs {
		result[i] = BytesToByteArray(common.Hash(acc.ABIEncode()))
	}
	return
}
