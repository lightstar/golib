#export TEST_CONFIG_ETCD_ENDPOINTS=127.0.0.1:2379

.PHONY: lint
lint:
	@golangci-lint run

.PHONY: test
test:
	@go test -race -cover -v ./...
