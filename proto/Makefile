help: ## Show this help
	@echo "Help"
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "    \033[36m%-20s\033[93m %s\n", $$1, $$2}'

move_to_master:
	mv ./gen/pb-go/* ../master/handlers/grpc/

generate-grpc: ## Generate the server and client from proto file
	rm -rf ./grpc/*.pb.go && \
	docker run \
	--rm -v ${PWD}:/defs namely/protoc-all:1.51_1 \
	-f ./*.proto \
	-l go \
	--go_out=plugins=grpc:/gen/grpc/ && \
	sleep 1 && \
	make move_to_master && \
	sleep 1 && \
	rm -rf ./gen