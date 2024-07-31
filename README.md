### Requisitos:
* Docker
  
### Utilização da biblioteca GoDotEnv para carregar todas as variáveis de ambiente de um arquivo .env
Endereço sobre a biblioteca [GoDotEnv](https://github.com/joho/godotenv)

Instalação:
```bash
$ go get github.com/joho/godotenv
```

### Utilização de uma biblioteca para geração de Logger
documentação [zap](https://pkg.go.dev/go.uber.org/zap#section-readme)

Instalação:
```bash
$ go get -u go.uber.org/zap
```
### utilização de uma biblioteca de conexão com SGBD MongoDB
Documentação [mongo](https://pkg.go.dev/go.mongodb.org/mongo-driver/mongo)

Instalação:
```bash
$ go get go.mongodb.org/mongo-driver/mongo
```
### pacote que gera e inspeciona UUIDs baseado no RFC 4122
Documentação [uuid](https://pkg.go.dev/github.com/google/uuid)

Instalação:
```bash
$ go get github.com/google/uuid
```



### iniciar um container Docker da imagem do banco de dados MongoDB com o nome de actionsDB

Subir o container:
```bash
$ docker container run -d -p 27017:27017 --name auctionsDB mongo
```

### inicializa o workspace
```bash
$ go work use .
```
Obs. Utilizei este comando porque ao executar o go run cmd/auction/main.go, apresentava erros de dependencias. Mesmo executando go mod tidy.
