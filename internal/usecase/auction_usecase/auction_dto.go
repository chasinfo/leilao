package auction_usecase

import "time"

type AuctionInputDTO struct {
	ProductName string              `json:"product_name"`
	Category    string              `json:"category"`
	Description string              `json:"description"`
	Condition   ProductionCondition `json:"condition"`
}

type AuctionOutputDTO struct {
	Id          string              `json:"id"`
	ProductName string              `json:"product_name"`
	Category    string              `json:"category"`
	Description string              `json:"description"`
	Condition   ProductionCondition `json:"condition"`
	Status      AuctionStatus       `json:"status"`
	Timestamp   time.Time           `json:"timestamp" time_format:"2006-01-02 15:04:05"`
}

type ProductionCondition int64
type AuctionStatus int64
