# gRPC-Mongo-Go

Showcase of gRPC, Mongo and GoLang

## File generation

### Protobuf

In project root run:

```shell
protoc -I .\gen\proto --go_out=. --go_opt=module=github.com/grpc-mongo-go --go-grpc_out=. --go-grpc_opt=module=github.com/grpc-mongo-go .\gen\proto\*.proto
```

### Mockery

Install mockery:

```shell
go install github.com/vektra/mockery/v2@v2.40.1
```

In project root run:

```shell
mockery
```

[Mockery documentation](https://vektra.github.io/mockery/latest/)

## Project start:

1) Start docker container using docker-compose.
2) Start Go server.
3) Start Go client and watch the result.
