# gRPC-Mongo-Go

## About

This project is a demonstration of the use of Go, gRPC, and MongoDB.

The gRPC server is responsible for:

- Adding a Blog record.
- Updating the Blog record.
- Deleting a Blog by ID.
- Returning a Blog by ID.
- Returning a list of all added Blogs.

The gRPC client is responsible for using the server and performing the following actions:

1) Create a new Blog record.
2) Retrieving the Blog by ID.
3) Update the Blog record.
4) Retrieving a list of all available Blogs.
5) Deleting a Blog by ID.

## Building

```shell
# build exe files.
make build
```

## Project startup

```shell
# build and start containers 
make build-start

# start containers
make start
```

## Requirements

### Protobuf

[Protobuf Installation guide](https://grpc.io/docs/protoc-installation/)

### Mockery

Install mockery:

```shell
go install github.com/vektra/mockery/v2@v2.40.1
```

[Mockery documentation](https://vektra.github.io/mockery/latest/)
