package poller

import (
	"context"
	"fmt"
	"time"

	"github.com/BOPR/common"
	"github.com/BOPR/config"
	"github.com/BOPR/core"
)

const (
	AggregatingService = "aggregator"
)

// Aggregator is the service which is supposed to create batches
// It has the following tasks:
// 1. Pick txs from the mempool
// 2. Validate these trnsactions
// 3. Update the DB post running each tx
// 4. Finally create a batch of all the transactions and post on-chain
type Aggregator struct {
	// Base service
	core.BaseService

	// DB instance
	DB core.DB

	// header listener subscription
	cancelAggregating context.CancelFunc
}

// NewAggregator returns new aggregator object
func NewAggregator(db core.DB) *Aggregator {
	// create logger
	logger := common.Logger.With("module", AggregatingService)
	aggregator := &Aggregator{}
	aggregator.BaseService = *core.NewBaseService(logger, AggregatingService, aggregator)
	aggregator.DB = db
	return aggregator
}

// OnStart starts new block subscription
func (a *Aggregator) OnStart() error {
	a.BaseService.OnStart() // Always call the overridden method.

	ctx, cancelAggregating := context.WithCancel(context.Background())
	a.cancelAggregating = cancelAggregating

	// start polling for checkpoint in buffer
	go a.startAggregating(ctx, config.GlobalCfg.PollingInterval)
	return nil
}

// OnStop stops all necessary go routines
func (a *Aggregator) OnStop() {
	a.BaseService.OnStop() // Always call the overridden method.

	// cancel ack process
	a.cancelAggregating()
}

func (a *Aggregator) startAggregating(ctx context.Context, interval time.Duration) {
	ticker := time.NewTicker(interval)
	// stop ticker when everything done
	defer ticker.Stop()
	for {
		select {
		case <-ticker.C:
			a.pickBatch()
			// pick batch from DB
		case <-ctx.Done():
			ticker.Stop()
			return
		}
	}
}

func (a *Aggregator) pickBatch() {
	txs, err := a.core.PopTxs()
	if err != nil {
		fmt.Println("Error while popping txs from mempool", "Error", err)
	}

	// Step-2
	// 1. Loop  all transactions
	// 2. Run verify_tx on them and update the leafs in DB according to result
	for i, tx := range txs {
		a.Logger.Debug("Verifing transaction", "index", i, "tx", tx.String())
		// Apply tx and get the updated accounts
		a.CheckTx(tx)

	}

	// Step-3
	// Finally create a merkel root of all updated leafs and push batch on-chain

}

// CheckTx fetches all the data required to validate tx from smart contact
// and calls the proccess tx function to return the updated balance root and accounts
func (a *Aggregator) CheckTx(tx core.Tx) {
	// fetch to account from DB
	// fromAccount, _ := a.core.GetAccount(tx.From)
	// fmt.Println("fetched account", fromAccount)

	// fromSiblings, err := core.FetchSiblings(fromAccount.Path, a.DB)
	// if err != nil {
	// 	fmt.Println("not able to fetch from siblings", "error", err)
	// }

	// // fetch from account from DB
	// toAccount, _ := a.core.GetAccount(tx.To)
	// fmt.Println("fetched account", toAccount)

	// toSiblings, err := core.FetchSiblings(toAccount.Path, a.DB)
	// if err != nil {
	// 	fmt.Println("not able to fetch to siblings", "error", err)
	// }

	// // fetch latest batch from DB
	// latestBatch, err := a.core.GetLatestBatch()
	// lbRootBytes, err := hex.DecodeString(latestBatch.StateRoot)
	// if err != nil {
	// 	fmt.Println("not able to fetch from siblings", "error", err)
	// }

	// newBalRoot, updatedFrom, updatedTo, err := core.ContractCallerObj.ProcessTx(core.BytesToByteArray(lbRootBytes),
	// 	tx, core.NewMerkleProof(fromAccount, fromSiblings),
	// 	core.NewMerkleProof(toAccount, toSiblings),
	// )
	// fmt.Println("all the updated data", newBalRoot, updatedFrom, updatedTo)
}
