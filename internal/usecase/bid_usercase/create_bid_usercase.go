package bid_usercase

import (
	"context"
	"os"
	"strconv"
	"time"

	"github.com/chasinfo/leilao/configuration/logger"
	"github.com/chasinfo/leilao/internal/entity/bid_entity"
	"github.com/chasinfo/leilao/internal/internal_error"
)

type BidUseCase struct {
	BidRepository       bid_entity.BidRepositoryInterface
	timer               *time.Timer
	maxBatchSize        int
	batchInsertInterval time.Duration
	bidChannel          chan bid_entity.Bid
}

var bidBatch []bid_entity.Bid

// constructor
func NewBidUseCase(bidRepository bid_entity.BidRepositoryInterface) BidUseCaseInterface {
	maxSizeInterval := getMaxBatchSizeInterval()
	maxBatchSize := getMaxBatchSize()

	bidUseCase := &BidUseCase{
		BidRepository:       bidRepository,
		maxBatchSize:        getMaxBatchSize(),
		batchInsertInterval: maxSizeInterval,
		timer:               time.NewTimer(maxSizeInterval),
		bidChannel:          make(chan bid_entity.Bid, maxBatchSize),
	}

	bidUseCase.triggerCreateRoutine(context.Background())

	return bidUseCase
}

func (bu *BidUseCase) triggerCreateRoutine(ctx context.Context) {
	go func() {
		defer close(bu.bidChannel)

		for {
			select {
			case bidEntity, ok := <-bu.bidChannel:
				if !ok {
					if len(bidBatch) > 0 {
						if err := bu.BidRepository.CreateBid(ctx, bidBatch); err != nil {
							logger.Error("Ocorreu um erro ao processar bit batch list.", err)
						}
					}
					return
				}
				bidBatch = append(bidBatch, bidEntity)

				if len(bidBatch) >= bu.maxBatchSize {
					if err := bu.BidRepository.CreateBid(ctx, bidBatch); err != nil {
						logger.Error("Ocorreu um erro ao processar bit batch list.", err)
					}

					bidBatch = nil
					bu.timer.Reset(bu.batchInsertInterval)
				}

			case <-bu.timer.C:
				if err := bu.BidRepository.CreateBid(ctx, bidBatch); err != nil {
					logger.Error("Ocorreu um erro ao processar bit batch list.", err)
				}

				bidBatch = nil
				bu.timer.Reset(bu.batchInsertInterval)
			}
		}
	}()
}

func (bu *BidUseCase) CreateBid(ctx context.Context, bidInput BidInputDTO) *internal_error.InternalError {

	bidEntity, err := bid_entity.CreateBid(bidInput.UserId, bidInput.AuctionId, bidInput.Amount)

	if err != nil {
		return err
	}

	bu.bidChannel <- *bidEntity

	return nil
}

func getMaxBatchSizeInterval() time.Duration {
	batchInsertInterval := os.Getenv("BATCH_INSERT_INTERVAL")
	duration, err := time.ParseDuration(batchInsertInterval)

	if err != nil {
		return 3 * time.Minute
	}

	return duration
}

func getMaxBatchSize() int {
	value, err := strconv.Atoi(os.Getenv("MAX_BATCH_SIZE"))

	if err != nil {
		return 5
	}

	return value
}
