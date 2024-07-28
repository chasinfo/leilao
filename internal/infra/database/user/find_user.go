package user

import (
	"context"
	"errors"
	"fmt"

	"github.com/chasinfo/leilao/configuration/logger"
	"github.com/chasinfo/leilao/internal/entity/user_entity"
	"github.com/chasinfo/leilao/internal/internal_error"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserEntityMongo struct {
	Id   string `bson:"_id"`
	Name string `bson:"name"`
}

type UserRepositoty struct {
	Collection *mongo.Collection
}

// constructor
func NewUserRepository(database *mongo.Database) *UserRepositoty {
	return &UserRepositoty{
		Collection: database.Collection("user"),
	}
}

func (ur *UserRepositoty) FindUserById(ctx context.Context, userId string) (*user_entity.User, *internal_error.InternalError) {
	filter := bson.M{"_id": userId}

	var userEntityMongo UserEntityMongo
	err := ur.Collection.FindOne(ctx, filter).Decode(&userEntityMongo)

	if err != nil {

		mensagem := "Ocorreu um erro ao buscar o usuário pelo Id"

		if errors.Is(err, mongo.ErrNoDocuments) {
			mensagem = fmt.Sprintf("Não foi encontrado nenhum usuário com este Id = %s", userId)

			logger.Error(mensagem, err)
			return nil, internal_error.NewNotFoundError(mensagem)
		}

		logger.Error(mensagem, err)
		return nil, internal_error.NewInternalServerError(mensagem)
	}

	userEntity := &user_entity.User{
		Id:   userEntityMongo.Id,
		Name: userEntityMongo.Name,
	}

	return userEntity, nil
}
