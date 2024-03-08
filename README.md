# todo-list-app

Esse projeto é um estudo para entender o funcionamento do gRPC como API no lado do servidor utilizando a linguagem Golang e o banco de dados sqlite3.

### Comando para gerar o protobufer da API:
```zsh
$ protoc --go_out=internal/pb --go_opt=paths=source_relative --go-grpc_out=internal/pb --go-grpc_opt=paths=source_relative api/api.proto
```

### Comando para rodar o projeto:
```zsh
$ go run ./cmd/server/main.go
```