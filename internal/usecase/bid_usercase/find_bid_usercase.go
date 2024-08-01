package bid_usercase

import (
	"context"

	"github.com/chasinfo/leilao/internal/internal_error"
)

func (bu *BidUseCase) FindBidByAuctionId(ctx context.Context, auctionId string) ([]BidOutputDTO, *internal_error.InternalError) {

	bidEntities, err := bu.FindBidByAuctionId(ctx, auctionId)

	if err != nil {
		return nil, err
	}

	var bidOutputs []BidOutputDTO

	for _, bidEntity := range bidEntities {
		bidOutputs = append(bidOutputs, BidOutputDTO{
			Id:        bidEntity.Id,
			UserId:    bidEntity.UserId,
			AuctionId: bidEntity.AuctionId,
			Amount:    bidEntity.Amount,
			Timestamp: bidEntity.Timestamp,
		})
	}

	return bidOutputs, nil
}

func (bu *BidUseCase) FindWinningBidByAuctionId(ctx context.Context, auctionId string) (*BidOutputDTO, *internal_error.InternalError) {

	bidEntity, err := bu.FindWinningBidByAuctionId(ctx, auctionId)

	if err != nil {
		return nil err
	}

	return &BidOutputDTO{
		Id:        bidEntity.Id,
		UserId:    bidEntity.UserId,
		AuctionId: bidEntity.AuctionId,
		Amount:    bidEntity.Amount,
		Timestamp: bidEntity.Timestamp,
	}, nil
}
