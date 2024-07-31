package auction_usecase

import (
	"context"

	"github.com/chasinfo/leilao/internal/entity/auction_entity"
	"github.com/chasinfo/leilao/internal/internal_error"
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
