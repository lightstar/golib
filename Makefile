#export TEST_CONFIG_ETCD_ENDPOINTS=127.0.0.1:2379
#export TEST_REDIS_ADDRESS=127.0.0.1:6379
#export TEST_MONGO_ADDRESS=127.0.0.1:27017

.PHONY: lint
lint:
	@golangci-lint run

.PHONY: test
test:
	@go test -race -cover -v ./...
