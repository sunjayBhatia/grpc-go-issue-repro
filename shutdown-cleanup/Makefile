TEST_COUNT ?= 1

.PHONY: generate-pb-go
generate-pb-go:
	protoc  --proto_path=${PWD} --go-grpc_out=. --go_out=. echo.proto

.PHONY: test
test:
	go test -v -count=${TEST_COUNT} .

.PHONY: test-race
test-race:
	go test -v -race -count=${TEST_COUNT} .
