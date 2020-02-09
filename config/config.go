package config

import (
	"crypto/ecdsa"
	"encoding/hex"
	"errors"
	"time"

	ethCmn "github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
)

const (
	DefaultMongoDB            = "mongodb://localhost:27017"
	DefaultEthRPC             = "http://localhost:8545"
	DefaultPollingInterval    = 5 * time.Second
	DefaultSeverPort          = "8080"
	DefaultConfirmationBlocks = 5
)

var GlobalCfg Configuration
var OperatorKey *ecdsa.PrivateKey
var OperatorPubKey *ecdsa.PublicKey

// Configuration represents heimdall config
type Configuration struct {
	MongoDB              string        `mapstructure:"mongo_DB_url"`
	EthRPC               string        `mapstructure:"eth_RPC_URL"`
	PollingInterval      time.Duration `mapstructure:"polling_interval"`
	ServerPort           string        `mapstructure:"server_port"`
	ConfirmationBlocks   uint64        `mapstructure:"confirmation_blocks"` // Number of blocks for confirmation
	RollupAddress        string        `mapstructure:"rollup_address"`
	MerkleTreeLibAddress string        `mapstructure:"merkle_lib_address"`
	OperatorKey          string        `mapstructure:"operator_key"`
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
		OperatorKey:          "",
	}
}

// SetOperatorKey sets the operatorKey globally
func SetOperatorKeys(privKeyStr string) error {
	privKeyBytes, err := hex.DecodeString(privKeyStr)
	if err != nil {
		return err
	}
	OperatorKey = crypto.ToECDSAUnsafe(privKeyBytes)
	publicKey := OperatorKey.Public()
	ecsdaPubKey, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		return errors.New("cannot assert type: publicKey is not of type *ecdsa.PublicKey")
	}
	OperatorPubKey = ecsdaPubKey
	return nil
}

func OperatorAddress() ethCmn.Address {
	address := crypto.PubkeyToAddress(*OperatorPubKey)
	return address
}
