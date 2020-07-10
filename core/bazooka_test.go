package core

import (
	"testing"

	"github.com/BOPR/config"
	"github.com/stretchr/testify/require"
)

func TestEncodeAccount(t *testing.T) {
	err := config.ParseAndInitGlobalConfig("../.")
	require.Equal(t, err, nil, "error should be nil")

	bazooka, err := NewPreLoadedBazooka()
	require.Equal(t, err, nil, "error should be nil")

	accountBytes, err := bazooka.EncodeAccount(1, 2, 3, 4)
	require.Equal(t, err, nil, "error should be nil")

	id, balance, nonce, token, err := bazooka.DecodeAccount(accountBytes)
	require.Equal(t, err, nil, "error should be nil")

	require.Equal(t, id.Uint64(), 1, "ID should match")
	require.Equal(t, balance.Uint64(), 2, "balance should match")
	require.Equal(t, nonce.Uint64(), 3, "nonce should match")
	require.Equal(t, token.Uint64(), 4, "token should match")
}

func TestEncodeTx(t *testing.T) {
	err := config.ParseAndInitGlobalConfig("../.")
	require.Equal(t, err, nil, "error should be nil")
	bazooka, err := NewPreLoadedBazooka()
	require.Equal(t, err, nil, "error should be nil")
	txBytes, err := bazooka.EncodeTx(1, 2, 3, 4, 5, 6)
	require.Equal(t, err, nil, "error should be nil")
	from, to, token, nonce, txType, amount, err := bazooka.DecodeTx(txBytes)
	require.Equal(t, err, nil, "error should be nil")
	require.Equal(t, from.Uint64(), uint64(1), "from should match")
	require.Equal(t, to.Uint64(), uint64(2), "to should match")
	require.Equal(t, token.Uint64(), uint64(3), "token should match")
	require.Equal(t, nonce.Uint64(), uint64(4), "nonce should match")
	require.Equal(t, amount.Uint64(), uint64(5), "amount should match")
	require.Equal(t, txType.Uint64(), uint64(6), "tx type should match")
}
