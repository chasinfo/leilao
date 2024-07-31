package auction

import (
	"context"

	"github.com/chasinfo/leilao/configuration/logger"
	"github.com/chasinfo/leilao/internal/entity/auction_entity"
	"github.com/chasinfo/leilao/internal/internal_error"
	"go.mongodb.org/mongo-driver/mongo"
)

type AuctionEntityMongo struct {
	Id          string                             `bson:"_id"`
	ProductName string                             `bson:"product_name"`
	Category    string                             `bson:"category"`
	Description string                             `bson:"description"`
	Condition   auction_entity.ProductionCondition `bson:"condition"`
	Status      auction_entity.AuctionStatus       `bson:"status"`
	Timestamp   int64                              `bson:"timestamp"`
}

type AuctionRepository struct {
	Collection *mongo.Collection
}

// construtor
func NewAuctionRepository(database *mongo.Database) *AuctionRepository {
	return &AuctionRepository{
		Collection: database.Collection("auctions"),
	}
}

func (ar *AuctionRepository) CreateAuction(ctx context.Context, auctionEntity *auction_entity.Auction) *internal_error.InternalError {

	AuctionEntityMongo := &AuctionEntityMongo{
		Id:          auctionEntity.Id,
		ProductName: auctionEntity.ProductName,
		Category:    auctionEntity.Category,
		Description: auctionEntity.Description,
		Condition:   auctionEntity.Condition,
		Status:      auctionEntity.Status,
		Timestamp:   auctionEntity.Timestamp.Unix(),
	}

	_, err := ar.Collection.InsertOne(ctx, AuctionEntityMongo)

	if err != nil {
		mensagem := "Ocorreu um erro ao tentar inserir auction"
		logger.Error(mensagem, err)
		return internal_error.NewInternalServerError(mensagem)
	}

	return nil
}
