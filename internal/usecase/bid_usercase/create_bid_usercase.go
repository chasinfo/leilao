package bid_usercase

import "github.com/chasinfo/leilao/internal/entity/bid_entity"

type BidUseCase struct {
	bidRepository bid_entity.BidRepositoryInterface
}
