package main

import (
	"context"
	"log"

	"github.com/chasinfo/leilao/tree/configuration/mongodb"
	"github.com/joho/godotenv"
)

func main() {
	ctx := context.Background()

	if err := godotenv.Load("cmd/auction/.env"); err != nil {
		log.Fatal("Ocorreu um erro ao carregar as variaveis env.")
		return
	}

	_, err := mongodb.NewMongoDBConnection(ctx)

	if err != nil {
		log.Fatal(err.Error())
		return
	}
}
