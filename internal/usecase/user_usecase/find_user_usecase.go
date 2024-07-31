package user_usecase

import (
	"context"

	"github.com/chasinfo/leilao/internal/entity/user_entity"
	"github.com/chasinfo/leilao/internal/internal_error"
)

type UserUseCase struct {
	userRepository user_entity.UserRepositoryInterface
}

func (u *UserUseCase) FindUserById(ctx context.Context, id string) (*UserOutputDTO, *internal_error.InternalError) {
	userEntity, err := u.userRepository.FindUserById(ctx, id)

	if err != nil {
		return nil, err
	}

	return &UserOutputDTO{
		Id:   userEntity.Id,
		Name: userEntity.Name,
	}, nil
}
