package auction_usecase

import (
	"context"

	"github.com/chasinfo/leilao/internal/entity/auction_entity"
	"github.com/chasinfo/leilao/internal/internal_error"
	"github.com/chasinfo/leilao/internal/usecase/bid_usercase"
)

func (au *AuctionUseCase) FindAuctionById(ctx context.Context, auctionId string) (*AuctionOutputDTO, *internal_error.InternalError) {

	auctionEntity, err := au.auctionRepository.FindAuctionById(ctx, auctionId)

	if err != nil {
		return nil, err
	}

	return &AuctionOutputDTO{
		Id:          auctionEntity.Id,
		ProductName: auctionEntity.ProductName,
		Category:    auctionEntity.Category,
		Description: auctionEntity.Description,
		Condition:   ProductionCondition(auctionEntity.Condition),
		Status:      AuctionStatus(auctionEntity.Status),
		Timestamp:   auctionEntity.Timestamp,
	}, nil
}

func (au *AuctionUseCase) FindAuctions(ctx context.Context, status AuctionStatus, category string, productName string) ([]AuctionOutputDTO, *internal_error.InternalError) {

	auctionEntities, err := au.auctionRepository.FindAuctions(ctx, auction_entity.AuctionStatus(status), category, productName)

	if err != nil {
		return nil, err
	}

	var auctionOutputs []AuctionOutputDTO

	for _, auctionEntity := range auctionEntities {
		auctionOutputs = append(auctionOutputs, AuctionOutputDTO{
			Id:          auctionEntity.Id,
			ProductName: auctionEntity.ProductName,
			Category:    auctionEntity.Category,
			Description: auctionEntity.Description,
			Condition:   ProductionCondition(auctionEntity.Condition),
			Status:      AuctionStatus(auctionEntity.Status),
			Timestamp:   auctionEntity.Timestamp,
		})
	}

	return auctionOutputs, nil
}

func (au *AuctionUseCase) FindWinningBidByAuctionId(ctx context.Context, auctionId string) (*WinningInfoOutputDTO, *internal_error.InternalError) {
	auction, err := au.auctionRepository.FindAuctionById(ctx, auctionId)

	if err != nil {
		return nil, err
	}

	auctionOutputDTO := AuctionOutputDTO{
		Id:          auction.Id,
		ProductName: auction.ProductName,
		Category:    auction.Category,
		Description: auction.Description,
		Condition:   ProductionCondition(auction.Condition),
		Status:      AuctionStatus(auction.Status),
		Timestamp:   auction.Timestamp,
	}

	bidWinning, err := au.bidRepository.FindWinningBidByAuctionId(ctx, auction.Id)

	if err != nil {
		return &WinningInfoOutputDTO{
			Auction: auctionOutputDTO,
			Bid:     nil,
		}, nil
	}

	BidOutputDTO := &bid_usercase.BidOutputDTO{
		Id:        bidWinning.Id,
		UserId:    bidWinning.UserId,
		AuctionId: bidWinning.AuctionId,
		Amount:    bidWinning.Amount,
		Timestamp: bidWinning.Timestamp,
	}

	return &WinningInfoOutputDTO{
		Auction: auctionOutputDTO,
		Bid:     BidOutputDTO,
	}, nil
}
