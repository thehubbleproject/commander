package core

import (
	"fmt"
	"math"
	"testing"

	"github.com/BOPR/common"
	"github.com/BOPR/core"

	"github.com/BOPR/aggregator"
	"github.com/BOPR/config"
	"github.com/stretchr/testify/require"
)

func TestPubkeyHashCreation(t *testing.T) {
	bz, err := core.ABIEncodePubkey("90718dcbc2477c86294742fb72bf098ba85ff671b88c8d79b2e09ce19bdbd88fd87047aaebc775b168372752aa8bc4e5be1ca5d39284fed00722f341927888c3")
	if err != nil {
		panic(err)
	}
	hash := common.Keccak256(bz)
	fmt.Println(hash.String())
}

func TestTxProcessing(t *testing.T) {
	err := config.ParseAndInitGlobalConfig("../.")
	require.Equal(t, err, nil, "Error while parsing and init config")
	db, err := core.NewDB("../.")
	require.Equal(t, err, nil, "Error while creating database")
	core.DBInstance = db

	bazooka, err := core.NewPreLoadedBazooka()
	require.Equal(t, err, nil, "Error while creating bazooka")
	core.LoadedBazooka = bazooka
	agg := aggregator.NewAggregator(db)
	genesisAccounts, err := core.LoadedBazooka.GetGenesisAccounts()
	require.Equal(t, err, nil, "error loading genesis accounts")
	zeroAccount := genesisAccounts[0]
	diff := int(math.Exp2(float64(4))) - len(genesisAccounts)
	var allAccounts []core.UserAccount

	// convert genesis accounts to user accounts
	for _, account := range genesisAccounts {
		pubkeyHash := core.ZERO_VALUE_LEAF.String()
		account.PublicKeyHash = pubkeyHash
		allAccounts = append(
			allAccounts,
			account,
		)
	}

	// fill the tree with zero leaves
	for diff > 0 {
		newAcc := core.EmptyAccount()
		newAcc.Data = zeroAccount.Data
		newAcc.Hash = core.ZERO_VALUE_LEAF.String()
		newAcc.PublicKeyHash = core.ZERO_VALUE_LEAF.String()
		allAccounts = append(allAccounts, newAcc)
		diff--
	}

	// load accounts
	err = core.DBInstance.InitBalancesTree(4, allAccounts)
	require.Equal(t, err, nil, "error initing balances tree")

	aliceAccountBytes, err := core.LoadedBazooka.EncodeAccount(2, 10, 0, 1)
	require.Equal(t, err, nil, "error encoding alice account")
	aliceAcc := core.NewUserAccount(2, core.STATUS_ACTIVE, "914873c8d5935837ade39cbdabd6efb3d3d4064c5918da11e555bba0ab2c58fee95974a3222830cf73d257bdc18cfcd01765482108a48e68bc0b657618acb40e", "0010", aliceAccountBytes)
	db.UpdateAccount(*aliceAcc)

	bobAccBytes, err := core.LoadedBazooka.EncodeAccount(3, 10, 0, 1)
	require.Equal(t, err, nil, "error encoding bob account")
	bobAcc := core.NewUserAccount(3, core.STATUS_ACTIVE, "90718dcbc2477c86294742fb72bf098ba85ff671b88c8d79b2e09ce19bdbd88fd87047aaebc775b168372752aa8bc4e5be1ca5d39284fed00722f341927888c3", "0011", bobAccBytes)
	db.UpdateAccount(*bobAcc)

	var txs []core.Tx
	txBytes, err := bazooka.EncodeTx(2, 3, 1, 1, 10, 1)
	if err != nil {
		panic(err)
	}
	tx1 := core.Tx{
		From: 2,
		To:   3,
		Data: txBytes,
	}
	for i := 0; i < 5; i++ {
		txs = append(txs, tx1)
	}
	err = agg.ProcessTx(txs)
	fmt.Println("err", err)

	require.Equal(t, err, nil, "error processing tx")
}
