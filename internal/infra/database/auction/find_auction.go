package auction

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/chasinfo/leilao/configuration/logger"
	"github.com/chasinfo/leilao/internal/entity/auction_entity"
	"github.com/chasinfo/leilao/internal/internal_error"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func (ar *AuctionRepository) FindAuctionById(ctx context.Context, auctionId string) (*auction_entity.Auction, *internal_error.InternalError) {

	filter := bson.M{"_id": auctionId}

	var auctionEntityMongo AuctionEntityMongo
	err := ar.Collection.FindOne(ctx, filter).Decode(&auctionEntityMongo)

	if err != nil {

		mensagem := "Ocorreu um erro ao buscar o auction pelo Id"

		if errors.Is(err, mongo.ErrNoDocuments) {
			mensagem = fmt.Sprintf("Não foi encontrado nenhum auction com este Id = %s", auctionId)

			logger.Error(mensagem, err)
			return nil, internal_error.NewNotFoundError(mensagem)
		}

		logger.Error(mensagem, err)
		return nil, internal_error.NewInternalServerError(mensagem)
	}

	auctionEntity := &auction_entity.Auction{
		Id:          auctionEntityMongo.Id,
		ProductName: auctionEntityMongo.ProductName,
		Category:    auctionEntityMongo.Category,
		Description: auctionEntityMongo.Description,
		Condition:   auctionEntityMongo.Condition,
		Status:      auctionEntityMongo.Status,
		Timestamp:   time.Unix(auctionEntityMongo.Timestamp, 0),
	}

	return auctionEntity, nil
}

func (ar *AuctionRepository) FindAuctions(
	ctx context.Context,
	status auction_entity.AuctionStatus,
	category string,
	productName string) ([]auction_entity.Auction, *internal_error.InternalError) {

	filter := bson.M{}

	if status != 0 {
		filter["status"] = status
	}

	if category != "" {
		filter["category"] = category
	}

	if productName != "" {
		filter["productName"] = primitive.Regex{
			Pattern: productName,
			Options: "i", // vai recuperar case sensitive: tanto maiúscula como minúscula.
		}
	}

	cursor, err := ar.Collection.Find(ctx, filter)

	if err != nil {

		mensagem := "Ocorreu um erro ao procurar auctions"
		logger.Error(mensagem, err)
		return nil, internal_error.NewInternalServerError(mensagem)
	}

	defer cursor.Close(ctx)

	var auctionEntityMongo []AuctionEntityMongo

	if err := cursor.All(ctx, &auctionEntityMongo); err != nil {
		mensagem := "Ocorreu um erro ao procurar auctions"
		logger.Error(mensagem, err)
		return nil, internal_error.NewInternalServerError(mensagem)
	}

	var auctionsEntity []auction_entity.Auction

	for _, auctionMongo := range auctionEntityMongo {
		auctionsEntity = append(auctionsEntity, auction_entity.Auction{
			Id:          auctionMongo.Id,
			ProductName: auctionMongo.ProductName,
			Category:    auctionMongo.Category,
			Description: auctionMongo.Description,
			Condition:   auctionMongo.Condition,
			Status:      auctionMongo.Status,
			Timestamp:   time.Unix(auctionMongo.Timestamp, 0),
		})
	}
	return auctionsEntity, nil
}
