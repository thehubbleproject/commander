package config

import (
	"time"

	ethCmn "github.com/ethereum/go-ethereum/common"
	"github.com/tendermint/tendermint/crypto/secp256k1"
)

const (
	DefaultMongoDB            = "mongodb://localhost:27017"
	DefaultEthRPC             = "http://localhost:8545"
	DefaultPollingInterval    = 5 * time.Second
	DefaultSeverPort          = "8080"
	DefaultConfirmationBlocks = 5
)

var GlobalCfg Configuration
var OperatorKey secp256k1.PrivKeySecp256k1

// Configuration represents heimdall config
type Configuration struct {
	MongoDB              string        `mapstructure:"mongo_DB_url"`
	EthRPC               string        `mapstructure:"eth_RPC_URL"`
	PollingInterval      time.Duration `mapstructure:"polling_interval"`
	ServerPort           string        `mapstructure:"server_port"`
	ConfirmationBlocks   uint64        `mapstructure:"confirmation_blocks"` // Number of blocks for confirmation
	RollupAddress        string        `mapstructure:"rollup_address"`
	MerkleTreeLibAddress string        `mapstructure:"merkle_lib_address"`
	OperatorPrivKey      string        `mapstructure:"operator_priv_key"`
}

// GetDefaultConfig returns the default configration options
func GetDefaultConfig() Configuration {
	return Configuration{
		MongoDB:              DefaultMongoDB,
		EthRPC:               DefaultEthRPC,
		PollingInterval:      DefaultPollingInterval,
		ServerPort:           DefaultSeverPort,
		ConfirmationBlocks:   DefaultConfirmationBlocks,
		RollupAddress:        ethCmn.Address{}.String(),
		MerkleTreeLibAddress: ethCmn.Address{}.String(),
	}
}
