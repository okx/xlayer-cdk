package bridgesync

import (
	"context"
	"time"

	"github.com/0xPolygon/cdk/etherman"
	"github.com/0xPolygon/cdk/sync"
	"github.com/ethereum/go-ethereum/common"
)

const (
	bridgeSyncL1       = "bridgesyncl1"
	bridgeSyncL2       = "bridgesyncl2"
	downloadBufferSize = 1000
)

var (
	retryAfterErrorPeriod      = time.Second * 10
	maxRetryAttemptsAfterError = 5
)

type LocalBridgeSync struct {
	processor *processor
	driver    *sync.EVMDriver
}

// NewL1 creates a bridge syncer that synchronizes the mainnet exit tree
func NewL1(
	ctx context.Context,
	dbPath string,
	bridge common.Address,
	syncBlockChunkSize uint64,
	blockFinalityType etherman.BlockNumberFinality,
	rd sync.ReorgDetector,
	ethClient EthClienter,
	initialBlock uint64,
) (*LocalBridgeSync, error) {
	return new(
		ctx,
		dbPath,
		bridge,
		syncBlockChunkSize,
		blockFinalityType,
		rd,
		ethClient,
		initialBlock,
		bridgeSyncL1,
	)
}

// NewL2 creates a bridge syncer that synchronizes the local exit tree
func NewL2(
	ctx context.Context,
	dbPath string,
	bridge common.Address,
	syncBlockChunkSize uint64,
	blockFinalityType etherman.BlockNumberFinality,
	rd sync.ReorgDetector,
	ethClient EthClienter,
	initialBlock uint64,
) (*LocalBridgeSync, error) {
	return new(
		ctx,
		dbPath,
		bridge,
		syncBlockChunkSize,
		blockFinalityType,
		rd,
		ethClient,
		initialBlock,
		bridgeSyncL2,
	)
}

func new(
	ctx context.Context,
	dbPath string,
	bridge common.Address,
	syncBlockChunkSize uint64,
	blockFinalityType etherman.BlockNumberFinality,
	rd sync.ReorgDetector,
	ethClient EthClienter,
	initialBlock uint64,
	l1OrL2ID string,
) (*LocalBridgeSync, error) {
	processor, err := newProcessor(ctx, dbPath, l1OrL2ID)
	if err != nil {
		return nil, err
	}
	lastProcessedBlock, err := processor.GetLastProcessedBlock(ctx)
	if err != nil {
		return nil, err
	}
	if lastProcessedBlock < initialBlock {
		err = processor.ProcessBlock(ctx, sync.Block{
			Num: initialBlock,
		})
		if err != nil {
			return nil, err
		}
	}
	rh := &sync.RetryHandler{
		MaxRetryAttemptsAfterError: maxRetryAttemptsAfterError,
		RetryAfterErrorPeriod:      retryAfterErrorPeriod,
	}

	appender, err := buildAppender(ethClient, bridge)
	if err != nil {
		return nil, err
	}
	downloader, err := sync.NewEVMDownloader(
		ethClient,
		syncBlockChunkSize,
		blockFinalityType,
		waitForNewBlocksPeriod,
		appender,
		[]common.Address{bridge},
		rh,
	)
	if err != nil {
		return nil, err
	}

	driver, err := sync.NewEVMDriver(rd, processor, downloader, l1OrL2ID, downloadBufferSize, rh)
	if err != nil {
		return nil, err
	}
	return &LocalBridgeSync{
		processor: processor,
		driver:    driver,
	}, nil
}

// Start starts the synchronization process
func (s *LocalBridgeSync) Start(ctx context.Context) {
	s.driver.Sync(ctx)
}