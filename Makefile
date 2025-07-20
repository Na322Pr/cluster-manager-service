LOCAL_BIN := $(shell go env GOPATH)/bin
PROTO_ROOT := api
GEN_DIR := pkg


run:
	go run cmd/main.go -config ./config/config.yaml

# proto generate

install:
	brew install bufbuild/buf/buf

bin-deps:
	GOBIN=$(LOCAL_BIN) go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
	GOBIN=$(LOCAL_BIN) go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest

generate: clean bin-deps
	PATH="$(LOCAL_BIN):$$PATH" buf generate

clean:
	rm -rf $(GEN_DIR)/api

# Метрики

run-prometheus:
	prometheus --config.file config/prometheus.yaml

run-grafana:
	brew services start grafana