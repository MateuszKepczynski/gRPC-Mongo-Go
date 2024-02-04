test:
	@echo "Running tests"
	@go test ./...

build-linux:
	@echo "Building Linux server executable"
	@go build -o bin/server/server ./cmd/server/
	@echo "Building Linux client executable"
	@go build -o bin/client/client ./cmd/client/

build:
	@echo "Building Windows server executable"
	@go build -o bin/server/server.exe ./cmd/server/
	@echo "Building Windows client executable"
	@go build -o bin/client/client.exe ./cmd/client/

build-start:
	@echo "Building and running app via docker-compose"
	@docker-compose -f ./docker/docker-compose.yaml up --build

start:
	@echo "Starting app via docker-compose"
	@docker-compose -f ./docker/docker-compose.yaml up

start-server:
	@echo "Opening Server Go Binary File"
	@./bin/server/server.exe

start-client:
	@echo "Opening Client Go Binary File"
	@./bin/client/client.exe

generate-proto:
	@echo "Generating Go Proto files"
	@protoc -I ./gen/proto --go_out=. --go_opt=module=github.com/grpc-mongo-go --go-grpc_out=. --go-grpc_opt=module=github.com/grpc-mongo-go ./gen/proto/*.proto

generate-mocks:
	@echo "Generating mock files"
	@mockery