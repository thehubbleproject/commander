package poller

import (
	"context"
	"fmt"
	"time"

	"github.com/BOPR/common"
	"github.com/BOPR/config"
	db "github.com/BOPR/db"
	"github.com/BOPR/types"

	tmCmn "github.com/tendermint/tendermint/libs/common"
)

const (
	AggregatingService = "aggregator"
)

type Aggregator struct {
	// Base service
	tmCmn.BaseService

	// DB instance
	DB db.DB

	// header listener subscription
	cancelAggregating context.CancelFunc
}

// NewAggregator returns new service object
func NewAggregator(db db.DB) *Aggregator {
	// create logger
	logger := common.Logger.With("module", AggregatingService)
	aggregator := &Aggregator{}
	aggregator.BaseService = *tmCmn.NewBaseService(logger, AggregatingService, aggregator)
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
	txs, err := a.DB.PopTxs()
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

// ApplyTx fetches all the data required to validate tx from smart contact
// and calls the proccess tx function to return the updated balance root and accounts
func (a *Aggregator) CheckTx(tx types.Tx) {
	// fetch to account from DB
	fromAccount, _ := a.DB.GetAccount(tx.From)
	fmt.Println("fetched account", fromAccount)

	fromSiblings, err := db.FetchSiblings(fromAccount.Path, a.DB)
	if err != nil {
		fmt.Println("not able to fetch from siblings", "error", err)
	}

	// fetch from account from DB
	toAccount, _ := a.DB.GetAccount(tx.To)
	fmt.Println("fetched account", toAccount)

	toSiblings, err := db.FetchSiblings(toAccount.Path, a.DB)
	if err != nil {
		fmt.Println("not able to fetch to siblings", "error", err)
	}

	// fetch latest batch from DB
	latestBatch, err := a.DB.GetLatestBatch()
	newBalRoot, updatedFrom, updatedTo, err := types.ContractCallerObj.ProcessTx(latestBatch.StateRoot,
		tx, types.NewMerkleProof(fromAccount, fromSiblings),
		types.NewMerkleProof(toAccount, toSiblings),
	)
	fmt.Println("all the updated data", newBalRoot, updatedFrom, updatedTo)
}
