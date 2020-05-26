package poller

import (
	"context"
	"fmt"
	"time"

	"github.com/BOPR/bazooka"
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

	// contract caller to interact with contracts
	loadedBazooka bazooka.Bazooka

	// DB instance
	DB core.DB

	// header listener subscription
	cancelAggregating context.CancelFunc
}

// NewAggregator returns new aggregator object
func NewAggregator(db core.DB) *Aggregator {
	// create logger
	logger := common.Logger.With("module", AggregatingService)
	loadedBazooka, err := bazooka.NewPreLoadedBazooka()
	if err != nil {
		panic(err)
	}
	aggregator := &Aggregator{}
	aggregator.BaseService = *core.NewBaseService(logger, AggregatingService, aggregator)
	aggregator.DB = db
	aggregator.loadedBazooka = loadedBazooka
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
	err = a.CheckTx(txs)
	if err != nil {
		fmt.Println("Error while processing tx", "error", err)
	}

	// Step-3
	// Finally create a merkel root of all updated leafs and push batch on-chain

}

// CheckTx fetches all the data required to validate tx from smart contact
// and calls the proccess tx function to return the updated balance root and accounts
func (a *Aggregator) CheckTx(txs []core.Tx) error {
	rootAcc, err := a.DB.GetRoot()
	if err != nil {
		return err
	}

	currentRoot, err := core.HexToByteArray(rootAcc.Hash)
	if err != nil {
		return err
	}

	for _, tx := range txs {
		fromAccProof, toAccProof, PDAproof, err := a.DB.GetTxVerificationData(tx)
		if err != nil {
			fmt.Println(err)
			return err
		}

		updatedRoot, updatedFromAcc, updatedToAcc, err := a.loadedBazooka.ProcessTx(currentRoot, tx, fromAccProof, toAccProof, PDAproof)
		if err != nil {
			fmt.Println(err)
			return err
		}

		err = a.DB.UpdateAccount(updatedFromAcc)
		if err != nil {
			return err
		}

		err = a.DB.UpdateAccount(updatedToAcc)
		if err != nil {
			return err
		}

		currentRoot = updatedRoot
	}

	return nil
}
