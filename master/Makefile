.PHONY: tests

help: ## Show this help
	@echo "Help"
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "    \033[36m%-20s\033[93m %s\n", $$1, $$2}'

tests: ## Run tests with race detection
	go test -race -v ./...

execute: ## Run the application
	go run ./cmd/main.go

install: ## Install dependencies
	go mod download && go mod tidy

db-run: ## Start the database container
	bash db-start.sh

run: install execute ## Run the microservice

generate: generate-grpc-v1 install ## Generate everything to run the app

dependency-injection: ## Generate the dependency injections using wire
	go generate -tags=wireinject ./...

generate-grpc-v1: ## Generate the server and client from proto file
	rm -rf ./pkg/adapters/grpc/v1/*.pb.go && \
	bash runtime-container.sh run \
	--rm -v ${PWD}:/defs namely/protoc-all \
	-f ./api_schemas/proto/v1/*.proto \
	-l go \
	--go_out=plugins=grpc:/gen/grpc/v1 && \
	sleep 1 && \
	mv ./gen/pb-go/* ./pkg/adapters/grpc/v1/ && \
	sleep 1 && \
	rm -rf ./gen
