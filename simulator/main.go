package simulator

import (
	"context"
	"time"

	"github.com/BOPR/common"
	"github.com/BOPR/core"
)

const (
	SimulatorService = "simulator"
)

type Simulator struct {
	// Base service
	core.BaseService

	// DB instance
	DB core.DB

	// contract caller to interact with contracts
	LoadedBazooka core.Bazooka

	// header listener subscription
	cancelSimulator context.CancelFunc

	toSwap bool
}

// NewSimulator returns new simulator object
func NewSimulator() *Simulator {
	logger := common.Logger.With("module", SimulatorService)
	sim := &Simulator{}
	sim.BaseService = *core.NewBaseService(logger, SimulatorService, sim)
	db, err := core.NewDB()
	if err != nil {
		panic(err)
	}

	sim.LoadedBazooka, err = core.NewPreLoadedBazooka()
	if err != nil {
		panic(err)
	}

	sim.DB = db
	return sim
}

// OnStart starts new block subscription
func (s *Simulator) OnStart() error {
	s.BaseService.OnStart() // Always call the overridden method.

	ctx, cancelSimulator := context.WithCancel(context.Background())
	s.cancelSimulator = cancelSimulator

	go s.SimulationStart(ctx, 10*time.Second)

	s.toSwap = false
	return nil
}

// OnStop stops all necessary go routines
func (s *Simulator) OnStop() {
	s.BaseService.OnStop() // Always call the overridden method.

	s.cancelSimulator()
}

// SimulationStart starts the simulator
func (s *Simulator) SimulationStart(ctx context.Context, interval time.Duration) {
	ticker := time.NewTicker(interval)
	// stop ticker when everything done
	defer ticker.Stop()
	for {
		select {
		case <-ticker.C:
			s.sendTxsToAndFro()
			// pick batch from DB
		case <-ctx.Done():
			ticker.Stop()
			return
		}
	}
}

// tries sending transactins to and fro accounts to the rollup node
func (s *Simulator) sendTxsToAndFro() {
	transferAmount := 1
	FromID := uint64(2)
	ToID := uint64(3)

	if s.toSwap {
		tempID := FromID
		FromID = ToID
		ToID = tempID
		s.toSwap = !s.toSwap
	}

	for i := 0; i < 7; i++ {
		latestFromAcc, err := s.DB.GetAccountByID(FromID)
		if err != nil {
			s.Logger.Error("unable to fetch latest account", "error", err)
			return
		}
		_, _, nonce, token, err := s.LoadedBazooka.DecodeAccount(latestFromAcc.Data)
		if err != nil {
			s.Logger.Error("unable to decode account", "error", err)
			return
		}
		txBytes, err := s.LoadedBazooka.EncodeTransferTx(int64(FromID), int64(ToID), int64(token.Uint64()), int64(nonce.Uint64())+1, int64(transferAmount), 1)
		if err != nil {
			s.Logger.Error("unable to encode tx", "error", err)
			return
		}
		txCore := core.NewPendingTx(FromID, ToID, core.TX_TRANSFER_TYPE, "0x1ad4773ace8ee65b8f1d94a3ca7adba51ee2ca0bdb550907715b3b65f1e3ad9f69e610383dc9ceb8a50c882da4b1b98b96500bdf308c1bdce2187cb23b7d736f1b", txBytes)

		err = s.DB.InsertTx(&txCore)
		if err != nil {
			s.Logger.Error("unable to insert tx", "error", err)
			return
		}
		s.Logger.Info("Sent a tx!", "TxHash", txCore.TxHash, "From", txCore.From, "To", txCore.To)

		if txCore.From == uint64(2) {
			s.toSwap = true
		}
	}
}
