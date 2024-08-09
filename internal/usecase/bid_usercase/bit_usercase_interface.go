package bid_usercase

import (
	"context"

	"github.com/chasinfo/leilao/internal/internal_error"
)

type BidUseCaseInterface interface {
	CreateBid(ctx context.Context, bidInput BidInputDTO) *internal_error.InternalError
	FindBidByAuctionId(ctx context.Context, auctionId string) ([]BidOutputDTO, *internal_error.InternalError)
	FindWinningBidByAuctionId(ctx context.Context, auctionId string) (*BidOutputDTO, *internal_error.InternalError)
}
