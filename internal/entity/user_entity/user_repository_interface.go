package user_entity

import (
	"context"

	"github.com/chasinfo/leilao/internal/internal_error"
)

type UserRepositoryInterface interface {
	FindUserById(ctx context.Context, userId string) (*User, *internal_error.InternalError)
}
