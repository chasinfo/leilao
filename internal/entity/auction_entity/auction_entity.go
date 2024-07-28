package auction_entity

import "time"

type Auction struct {
	Id          string
	ProductName string
	Category    string
	Description string
	Condition   ProductionCondition
	Status      AuctionStatus
	Timestamp   time.Time
}

type ProductionCondition int
type AuctionStatus int

const (
	New         ProductionCondition = iota
	Used                            = 1
	Refurbished                     = 2
)

const (
	Active    AuctionStatus = iota
	Completed               = 1
)
