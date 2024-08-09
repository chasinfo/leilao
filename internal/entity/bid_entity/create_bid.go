package bid_entity

import (
	"time"

	"github.com/chasinfo/leilao/internal/internal_error"
	"github.com/google/uuid"
)

func CreateBid(userId, auctionId string, amout float64) (*Bid, *internal_error.InternalError) {
	bid := &Bid{
		Id:        uuid.New().String(),
		UserId:    userId,
		AuctionId: auctionId,
		Amount:    amout,
		Timestamp: time.Now(),
	}

	err := bid.Validate()

	if err != nil {
		return nil, err
	}

	return bid, nil
}

func (b *Bid) Validate() *internal_error.InternalError {

	err := uuid.Validate(b.UserId)
	if err != nil {
		return internal_error.NewBadRequestError("UserId não é um id válido.")
	}

	if err := uuid.Validate(b.AuctionId); err != nil {
		return internal_error.NewBadRequestError("UserId não é um id válido.")
	}

	if b.Amount <= 0 {
		return internal_error.NewBadRequestError("A quantia não é um valor válido.")
	}

	return nil
}
