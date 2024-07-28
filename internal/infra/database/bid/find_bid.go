package bid

import (
	"context"
	"fmt"
	"time"

	"github.com/chasinfo/leilao/configuration/logger"
	"github.com/chasinfo/leilao/internal/entity/bid_entity"
	"github.com/chasinfo/leilao/internal/internal_error"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func (bd *BidRepository) FindBidByAuctionId(ctx context.Context, auctionId string) ([]bid_entity.Bid, *internal_error.InternalError) {
	filter := bson.M{"auction_id": auctionId}

	cursor, err := bd.Collection.Find(ctx, filter)

	if err != nil {
		mensagem := fmt.Sprintf("Ocorreu um erro na tentative de pesquisar por auctionId %s", auctionId)
		logger.Error(mensagem, err)
		return nil, internal_error.NewInternalServerError(mensagem)
	}

	defer cursor.Close(ctx)

	var bidEntityMongo []BidEntityMongo

	if err := cursor.All(ctx, &bidEntityMongo); err != nil {
		mensagem := "Ocorreu um erro ao procurar bids"
		logger.Error(mensagem, err)
		return nil, internal_error.NewInternalServerError(mensagem)
	}

	var bidEntities []bid_entity.Bid

	for _, bidMongo := range bidEntityMongo {
		bidEntities = append(bidEntities, bid_entity.Bid{
			Id:        bidMongo.Id,
			UserId:    bidMongo.UserId,
			AuctionId: bidMongo.AuctionId,
			Amount:    bidMongo.Amount,
			Timestamp: time.Unix(bidMongo.Timestamp, 0),
		})
	}

	return bidEntities, nil
}

func (bd *BidRepository) FindWinningBidByAuctionId(ctx context.Context, auctionId string) (*bid_entity.Bid, *internal_error.InternalError) {

	filter := bson.M{"auction_id": auctionId}
	opts := options.FindOne().SetSort(bson.D{{Key: "amount", Value: -1}})

	var bidEntityMongo BidEntityMongo
	if err := bd.Collection.FindOne(ctx, filter, opts).Decode(&bidEntityMongo); err != nil {
		mensagem := "Ocorreu um erro ao buscar o Bid"
		logger.Error(mensagem, err)
		return nil, internal_error.NewInternalServerError(mensagem)
	}

	return &bid_entity.Bid{
		Id:        bidEntityMongo.Id,
		UserId:    bidEntityMongo.UserId,
		AuctionId: bidEntityMongo.AuctionId,
		Amount:    bidEntityMongo.Amount,
		Timestamp: time.Unix(bidEntityMongo.Timestamp, 0),
	}, nil
}
