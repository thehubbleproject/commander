package types

import "math/big"

// Params stores all the parameters which are maintained on-chain and keeps updating
// them whenever they change on-chain
type Params struct {
	DBModel

	// Stake amount which coordinator needs to submit a new batch
	// Updates when syncer receives a stake update event from the contract
	// Used while sending new batch
	StakeAmount uint64 `json:"stakeAmount"`

	// MaxDepth is the maximum depth of the balances tree possible
	// If in case we want to increase it we will update the value on the contract
	// And then this will be updated
	MaxDepth uint64 `json:"maxDepth"`

	// DepositSubTreeHeight is the maximum height of the deposit subtree that the coordinator wants to merge
	// It is set on the contract and will be updated when that value changes
	MaxDepositSubTreeHeight uint64 `json:"maxDepositSubTreeHeight"`
}

// Maintains sync information
type SyncStatus struct {
	DBModel
	// Last eth block seen by the syncer is persisted here so that we can resume sync from it
	LastEthBlockRecorded uint64 `json:"lastEthBlockRecorded"`

	// Last batch index is recorded for this field
	LastBatchRecorded uint64 `json:"lastBatchRecorded"`
}

func (ss *SyncStatus) LastEthBlockBigInt() *big.Int {
	n := new(big.Int)
	return n.SetUint64(ss.LastEthBlockRecorded)

}
