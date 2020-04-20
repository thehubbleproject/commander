package types

import (
	"math/big"
	"testing"

	"github.com/BOPR/common"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/stretchr/testify/require"
)

func TestEmptyAccount(t *testing.T) {
	// create an empty inactive account
	acc, err := ABIEncode()
	require.NoError(t, err)
	require.Equal(t, ZERO_VALUE_LEAF.String(), common.Keccak256(acc).Hex()[2:], "The root leaves should match")
}

func ABIEncode() ([]byte, error) {
	uint256Ty, err := abi.NewType("uint256", "uint256", nil)
	if err != nil {
		return []byte(""), err
	}

	arguments := abi.Arguments{
		{
			Type: uint256Ty,
		},
	}
	bytes, err := arguments.Pack(
		big.NewInt(int64(0)),
	)
	if err != nil {
		return []byte(""), err
	}

	return bytes, nil
}
