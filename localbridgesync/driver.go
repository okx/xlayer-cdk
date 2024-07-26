package localbridgesync

import (
	"context"

	"github.com/0xPolygon/cdk/log"
	"github.com/0xPolygon/cdk/reorgdetector"
	"github.com/ethereum/go-ethereum/common"
)

const (
	downloadBufferSize = 1000
	reorgDetectorID    = "localbridgesync"
)

type downloaderFull interface {
	downloaderInterface
	download(ctx context.Context, fromBlock uint64, downloadedCh chan block)
}

type driver struct {
	reorgDetector ReorgDetector
	reorgSub      *reorgdetector.Subscription
	processor     processorInterface
	downloader    downloaderFull
}

type processorInterface interface {
	getLastProcessedBlock(ctx context.Context) (uint64, error)
	storeBridgeEvents(blockNum uint64, events []BridgeEvent) error
	reorg(firstReorgedBlock uint64) error
}

type ReorgDetector interface {
	Subscribe(id string) *reorgdetector.Subscription
	AddBlockToTrack(ctx context.Context, id string, blockNum uint64, blockHash common.Hash) error
}

func newDriver(
	reorgDetector ReorgDetector,
	processor processorInterface,
	downloader downloaderFull,
) (*driver, error) {
	reorgSub := reorgDetector.Subscribe(reorgDetectorID)
	return &driver{
		reorgDetector: reorgDetector,
		reorgSub:      reorgSub,
		processor:     processor,
		downloader:    downloader,
	}, nil
}

func (d *driver) Sync(ctx context.Context) {
reset:
	var (
		lastProcessedBlock uint64
		attempts           int
		err                error
	)
	for {
		lastProcessedBlock, err = d.processor.getLastProcessedBlock(ctx)
		if err != nil {
			attempts++
			log.Error("error geting last processed block: ", err)
			retryHandler("Sync", attempts)
			continue
		}
		break
	}
	cancellableCtx, cancel := context.WithCancel(ctx)
	defer cancel()

	// start downloading
	downloadCh := make(chan block, downloadBufferSize)
	go d.downloader.download(cancellableCtx, lastProcessedBlock, downloadCh)

	for {
		select {
		case b := <-downloadCh:
			log.Debug("handleNewBlock")
			d.handleNewBlock(ctx, b)
		case firstReorgedBlock := <-d.reorgSub.FirstReorgedBlock:
			log.Debug("handleReorg")
			d.handleReorg(cancel, downloadCh, firstReorgedBlock)
			goto reset
		}
	}
}

func (d *driver) handleNewBlock(ctx context.Context, b block) {
	attempts := 0
	for {
		err := d.reorgDetector.AddBlockToTrack(ctx, reorgDetectorID, b.Num, b.Hash)
		if err != nil {
			attempts++
			log.Errorf("error adding block %d to tracker: %v", b.Num, err)
			retryHandler("handleNewBlock", attempts)
			continue
		}
		break
	}
	attempts = 0
	for {
		err := d.processor.storeBridgeEvents(b.Num, b.Events)
		if err != nil {
			attempts++
			log.Errorf("error processing events for blcok %d, err: ", b.Num, err)
			retryHandler("handleNewBlock", attempts)
			continue
		}
		break
	}
}

func (d *driver) handleReorg(
	cancel context.CancelFunc, downloadCh chan block, firstReorgedBlock uint64,
) {
	// stop downloader
	cancel()
	_, ok := <-downloadCh
	for ok {
		_, ok = <-downloadCh
	}
	// handle reorg
	attempts := 0
	for {
		err := d.processor.reorg(firstReorgedBlock)
		if err != nil {
			attempts++
			log.Errorf(
				"error processing reorg, last valid block %d, err: %v",
				firstReorgedBlock, err,
			)
			retryHandler("handleReorg", attempts)
			continue
		}
		break
	}
	d.reorgSub.ReorgProcessed <- true
}