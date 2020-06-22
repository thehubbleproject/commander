package core

import (
	"fmt"
	"math"
	"math/big"
	"testing"

	"github.com/BOPR/common"
	"github.com/BOPR/config"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/stretchr/testify/require"
)

func TestEmptyAccount(t *testing.T) {
	// create an empty inactive account
	acc, err := ABIEncode()
	require.NoError(t, err)
	require.Equal(t, ZERO_VALUE_LEAF.String(), common.Keccak256(acc).Hex(), "The root leaves should match")
}

func TestNonEmptyAccount(t *testing.T) {
	acc := NewPendingUserAccount(1, 10, 1, "0x914873c8d5935837ade39cbdabd6efb3d3d4064c5918da11e555bba0ab2c58fee95974a3222830cf73d257bdc18cfcd01765482108a48e68bc0b657618acb40e")
	acc.CreateAccountHash()
	fmt.Println("account hash", acc.Hash, "0x94528175e1bb91e6e11ae8b0dae313996291b28a716b004f2ba0e5b3b6f64b23")
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

func TestPubkeyHashCreation(t *testing.T) {
	bz, err := ABIEncodePubkey("90718dcbc2477c86294742fb72bf098ba85ff671b88c8d79b2e09ce19bdbd88fd87047aaebc775b168372752aa8bc4e5be1ca5d39284fed00722f341927888c3")
	if err != nil {
		panic(err)
	}

	hash := common.Keccak256(bz)

	fmt.Println(hash.String())
}

func TestInitBalanceTree(t *testing.T) {
	db, err := NewDB()
	if err != nil {
		panic(err)
	}
	var genesisAccounts []UserAccount

	genesis, err := config.ReadGenesisFile()
	common.PanicIfError(err)

	err = genesis.Validate()
	if err != nil {
		common.PanicIfError(err)
	}

	diff := int(math.Exp2(float64(genesis.MaxTreeDepth))) - len(genesis.GenesisAccounts.Accounts)
	var allAccounts []UserAccount

	// convert genesis accounts to user accounts
	for _, account := range genesis.GenesisAccounts.Accounts {
		pubkeyHash := ZERO_VALUE_LEAF.String()
		allAccounts = append(
			allAccounts,
			UserAccount{
				AccountID:     account.ID,
				Balance:       account.Balance,
				TokenType:     account.TokenType,
				Nonce:         account.Nonce,
				Status:        account.Status,
				PublicKey:     account.PublicKey,
				PublicKeyHash: pubkeyHash,
			},
		)
	}

	// fill the tree with zero leaves
	for diff > 0 {
		newAcc := EmptyAccount()
		newAcc.Hash = ZERO_VALUE_LEAF.String()
		newAcc.PublicKeyHash = ZERO_VALUE_LEAF.String()
		allAccounts = append(allAccounts, newAcc)
		diff--
	}
	err = db.InitBalancesTree(2, genesisAccounts)
	require.Nil(t, err, "Error creating init balance tree")
}
