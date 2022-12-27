help: ## Show this help
	@echo "Help"
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "    \033[36m%-20s\033[93m %s\n", $$1, $$2}'

move_to_master:
	mv ./gen/pb-go/* ../master/handlers/grpc/

generate-grpc: ## Generate the server and client from proto file
	bash build_grpc.sh & \
	sleep 15 && \
	make move_to_master && \
	sleep 1 && \
	rm -rf ./gen && \
	ps | grep -h docker | grep -v grep | awk -F ' ' '{print $$1}' | while read line; do kill -9 "$${line}" 2>/dev/null 2>&1; done && \
	killall Docker && sleep 20 && open -a Docker
