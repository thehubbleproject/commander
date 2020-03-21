package types

import (
	"math/big"

	"github.com/BOPR/common"
	"github.com/BOPR/contracts/rollup"
	"github.com/ethereum/go-ethereum/accounts/abi"
)

// UserAccount is the user data stored on the node per user
type UserAccount struct {
	// ID is the path of the user account in the PDA Tree
	// Cannot be changed once created
	ID uint64
	// Token type of the user account
	// Cannot be changed after creation
	TokenType uint64

	// Balance of the user account
	Balance uint64
	// Nonce of the account
	Nonce uint64
}

func NewUserAccount(id, balance, tokenType, nonce uint64) UserAccount {
	return UserAccount{
		ID:        id,
		Balance:   balance,
		TokenType: tokenType,
		Nonce:     nonce,
	}
}

func (acc *UserAccount) ToABIAccount() rollup.DataTypesAccount {
	return rollup.DataTypesAccount{
		ID:        UintToBigInt(acc.id),
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
