package types

// Params stores all the parameters which are maintained on-chain and keeps updating
// them whenever they change on-chain
type Params struct {
	DBModel

	// Stake amount which coordinator needs to submit a new batch
	// Updates when syncer receives a stake update event from the contract
	// Used while sending new batch
	StakeAmount uint `json:"stakeAmount"`

	// MaxDepth is the maximum depth of the balances tree possible
	// If in case we want to increase it we will update the value on the contract
	// And then this will be updated
	MaxDepth uint `json:"maxDepth"`

	// DepositSubTreeHeight is the maximum height of the deposit subtree that the coordinator wants to merge
	// It is set on the contract and will be updated when that value changes
	MaxDepositSubTreeHeight uint `json:"maxDepositSubTreeHeight"`
}

// Maintains sync information
type SyncStatus struct {
	DBModel
	// Last eth block seen by the syncer is persisted here so that we can resume sync from it
	LastEthBlockRecorded uint `json:"lastEthBlockRecorded"`

	// Last batch index is recorded for this field
	LastBatchRecorded uint `json:"lastBatchRecorded"`
}

// func (l *ListenerLog) BigInt() *big.Int {
// 	n := new(big.Int)
// 	n, _ = n.SetString(l.LastRecordedBlock, 10)
// 	return n
// }
