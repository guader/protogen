.PHONY: build
build:
		@go build -o ./debug/bin/protoc-gen-go-err ./cmd/protoc-gen-go-err

.PHONY: debug
debug:
		@protoc -I ./debug/pb \
				--go_out ./debug/pb --go_opt paths=source_relative \
				--go-err_out ./debug/pb --go-err_opt paths=source_relative \
				./debug/pb/*.proto

.PHONY: all
all: build debug
