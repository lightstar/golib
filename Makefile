#export TEST_CONFIG_ETCD_ENDPOINTS=127.0.0.1:2379
#export TEST_REDIS_ADDRESS=127.0.0.1:6379
#export TEST_MONGO_ADDRESS=127.0.0.1:27017

.PHONY: lint
lint:
	@golangci-lint run

.PHONY: test
test:
	@go test -race -cover -v ./...

.PHONY: testproto
testproto:
	@protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. \
		--go-grpc_opt=paths=source_relative api/testproto/test.proto
