prepare:
	go1.20 get ./...

.PHONY: test
test:prepare
	go test -race -timeout 30s -cover ./...

.PHONY: lint
lint:prepare
	golangci-lint --timeout 5m run --out-format checkstyle ./...