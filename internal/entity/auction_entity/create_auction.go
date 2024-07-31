package auction_entity

import (
	"time"

	"github.com/chasinfo/leilao/internal/internal_error"
	"github.com/google/uuid"
)

func CreateAuction(
	productName string,
	category string,
	description string,
	condition ProductionCondition,
) (*Auction, *internal_error.InternalError) {
	auction := &Auction{
		Id:          uuid.New().String(),
		ProductName: productName,
		Category:    category,
		Description: description,
		Condition:   condition,
		Status:      Active,
		Timestamp:   time.Now(),
	}

	if err := auction.Validate(); err != nil {
		return nil, err
	}

	return auction, nil
}

func (au *Auction) Validate() *internal_error.InternalError {
	if len(au.ProductName) <= 1 ||
		len(au.Category) <= 2 ||
		len(au.Description) <= 10 &&
			(au.Condition != New) &&
			(au.Condition != Refurbished) &&
			(au.Condition != Used) {
		return internal_error.NewBadRequestError("Dados o Auction não são válidos")
	}

	return nil
}
