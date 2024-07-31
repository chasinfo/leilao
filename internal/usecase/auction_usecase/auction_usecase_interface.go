package auction_usecase

import (
	"context"

	"github.com/chasinfo/leilao/internal/internal_error"
)

type AuctionUseCaseInterface interface {
	CreateAuction(ctx context.Context, auctionInput AuctionInputDTO) *internal_error.InternalError
	FindAuctionById(ctx context.Context, auctionId string) (*AuctionOutputDTO, *internal_error.InternalError)
	FindAuctions(ctx context.Context, status AuctionStatus, category string, productName string) ([]AuctionOutputDTO, *internal_error.InternalError)
}
