package listener

import (
	"context"
	"fmt"
	"time"

	"github.com/BOPR/common"
	"github.com/BOPR/db"
	"github.com/BOPR/types"
	"github.com/ethereum/go-ethereum"
	ethTypes "github.com/ethereum/go-ethereum/core/types"

	"github.com/ethereum/go-ethereum/accounts/abi"
)

// Syncer to sync events from ethereum chain
type Syncer struct {
	// Base service
	BaseService

	// ABIs
	abis []*abi.ABI

	// storage client
	storageClient db.DB

	contractCaller types.ContractCaller

	// header channel
	HeaderChannel chan *ethTypes.Header
	// cancel function for poll/subscription
	cancelSubscription context.CancelFunc

	// header listener subscription
	cancelHeaderProcess context.CancelFunc
}

func NewSyncer() Syncer {
	// create logger
	logger := common.Logger.With(SyncerServiceName)

	// create syncer obj
	syncerService := &Syncer{}

	// create new base service
	syncerService.BaseService = *NewBaseService(logger, SyncerServiceName, syncerService)

	contractCaller, err := types.NewContractCaller()
	if err != nil {
		panic(err)
	}
	abis := []*abi.ABI{
		&contractCaller.RollupContractABI,
		&contractCaller.MerkleTreeABI,
	}
	syncerService.abis = abis

	return *syncerService
}

// OnStart starts new block subscription
func (s *Syncer) OnStart() error {
	s.BaseService.OnStart() // Always call the overridden method.

	// create cancellable context
	ctx, cancelSubscription := context.WithCancel(context.Background())
	s.cancelSubscription = cancelSubscription

	// create cancellable context
	headerCtx, cancelHeaderProcess := context.WithCancel(context.Background())
	s.cancelHeaderProcess = cancelHeaderProcess

	// start header process
	go s.startHeaderProcess(headerCtx)
	fmt.Println("ethclient", s.contractCaller.EthClient)
	// subscribe to new head
	subscription, err := s.contractCaller.EthClient.SubscribeNewHead(ctx, s.HeaderChannel)
	if err != nil {
		// start go routine to poll for new header using client object
		go s.startPolling(ctx, 1*time.Second)
	} else {
		// start go routine to listen new header using subscription
		go s.startSubscription(ctx, subscription)
	}

	return nil
}

// OnStop stops all necessary go routines
func (s *Syncer) OnStop() {
	s.BaseService.OnStop() // Always call the overridden method.

	// cancel subscription if any
	s.cancelSubscription()

	// cancel header process
	s.cancelHeaderProcess()
}

// startHeaderProcess starts header process when they get new header
func (s *Syncer) startHeaderProcess(ctx context.Context) {
	for {
		select {
		case newHeader := <-s.HeaderChannel:
			fmt.Println("new header found", newHeader)
			// s.processHeader(newHeader)
		case <-ctx.Done():
			return
		}
	}
}

// startPolling starts polling
func (s *Syncer) startPolling(ctx context.Context, pollInterval time.Duration) {
	// How often to fire the passed in function in second
	interval := pollInterval

	// Setup the ticket and the channel to signal
	// the ending of the interval
	ticker := time.NewTicker(interval)

	// start listening
	for {
		select {
		case <-ticker.C:
			header, err := s.contractCaller.EthClient.HeaderByNumber(ctx, nil)
			if err == nil && header != nil {
				fmt.Println("found new header", header)
				// send data to channel
				s.HeaderChannel <- header
			}
		case <-ctx.Done():
			ticker.Stop()
			return
		}
	}
}

func (s *Syncer) startSubscription(ctx context.Context, subscription ethereum.Subscription) {
	for {
		select {
		case err := <-subscription.Err():
			// stop service
			s.Logger.Error("Error while subscribing new blocks", "error", err)
			s.Stop()

			// cancel subscription
			s.cancelSubscription()
			return
		case <-ctx.Done():
			return
		}
	}
}
