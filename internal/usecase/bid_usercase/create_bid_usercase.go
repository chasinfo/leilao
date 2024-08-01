package bid_usercase

import "github.com/chasinfo/leilao/internal/entity/bid_entity"

type BidUseCase struct {
	BidRepository bid_entity.BidRepositoryInterface
}
