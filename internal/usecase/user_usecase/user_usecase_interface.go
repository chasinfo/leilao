package user_usecase

import (
	"context"

	"github.com/chasinfo/leilao/internal/internal_error"
)

type UserUseCaseInterface interface {
	FindUserById(ctx context.Context, id string) (*UserOutputDTO, *internal_error.InternalError)
}
