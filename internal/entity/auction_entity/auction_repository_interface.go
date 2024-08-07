package auction_entity

import (
	"context"

	"github.com/chasinfo/leilao/internal/internal_error"
)

type AuctionRepositoryInterface interface {
	CreateAuction(ctx context.Context, auctionEntity *Auction) *internal_error.InternalError
	FindAuctionById(ctx context.Context, auctionId string) (*Auction, *internal_error.InternalError)
	FindAuctions(ctx context.Context, status AuctionStatus, category string, productName string) ([]Auction, *internal_error.InternalError)
}
