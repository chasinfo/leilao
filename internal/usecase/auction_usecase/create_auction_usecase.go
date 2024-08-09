package auction_usecase

import (
	"context"

	"github.com/chasinfo/leilao/internal/entity/auction_entity"
	"github.com/chasinfo/leilao/internal/entity/bid_entity"
	"github.com/chasinfo/leilao/internal/internal_error"
)

type AuctionUseCase struct {
	auctionRepository auction_entity.AuctionRepositoryInterface
	bidRepository     bid_entity.BidRepositoryInterface
}

func (au *AuctionUseCase) CreateAuction(ctx context.Context, auctionInput AuctionInputDTO) *internal_error.InternalError {
	auction, err := auction_entity.CreateAuction(
		auctionInput.ProductName,
		auctionInput.Category,
		auctionInput.Description,
		auction_entity.ProductionCondition(auctionInput.Condition))

	if err != nil {
		return err
	}

	if err := au.auctionRepository.CreateAuction(ctx, auction); err != nil {
		return err
	}

	return nil
}
