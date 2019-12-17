package poller

import (
	"context"
	"fmt"
	"time"

	"github.com/BOPR/common"
	db "github.com/BOPR/db"

	tmCmn "github.com/tendermint/tendermint/libs/common"
)

const (
	AggregatingService = "aggregator"
)

type Aggregator struct {
	// Base service
	tmCmn.BaseService

	// header listener subscription
	cancelAggregating context.CancelFunc
}

// NewAggregator returns new service object
func NewAggregator() *Aggregator {
	// create logger
	logger := common.Logger.With("module", AggregatingService)
	aggregator := &Aggregator{}
	aggregator.BaseService = *tmCmn.NewBaseService(logger, AggregatingService, aggregator)
	return aggregator
}

// OnStart starts new block subscription
func (a *Aggregator) OnStart() error {
	a.BaseService.OnStart() // Always call the overridden method.

	ctx, cancelAggregating := context.WithCancel(context.Background())
	a.cancelAggregating = cancelAggregating

	// start polling for checkpoint in buffer
	go a.startAggregating(ctx, 5*time.Second)
	a.Logger.Info("Starting aggregator")
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
	txs, err := db.PopTxs()
	if err != nil {
		fmt.Println("Error while popping txs from mempool", "Error", err)
	}

	// Step-2
	// 1. Loop  all transactions
	// 2. Run verify_tx on them and update the leafs in DB according to result
	for i, tx := range txs {
		a.Logger.Debug("Verifing transaction", "index", i, "tx", tx.String())
		// TODO Run verify tx from contract

	}

	// Step-3
	// Finally create a merkel root of all updated leafs and push batch on-chain

}
