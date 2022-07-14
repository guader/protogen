# Protogen

Generate golang code from proto files.

Depends on code generated from `protoc-gen-go`.

## protoc-gen-go-enum

Generate code for enums.

### Usage

```bash
go get -u github.com/guader/protogen

go install github.com/guader/protogen/protoc-gen-go-enum

# example:
#  generate code for enum A and B with specified plugins.
# output files:
#  ./pb/*.pb.enum.err.go
#  ./pb/*.pb.enum.grpc_status.go
#  ./pb/*.pb.enum.number.go
#  ./pb/*.pb.enum.validation.go
protoc -I ./pb \
	--go_out ./pb \
	--go_opt paths=source_relative \
	--go-enum_out ./pb \
	--go-enum_opt paths=source_relative \
	--go-enum_opt plugins=err+grpcStatus+number+validation \
	--go-enum_opt enums=A+B \
	./pb/*.proto
```

### Options

- `plugins`: required, a list of plugins split by `+`, generate nothing if not specified.
    - `err`: implement `error` interface for enum, `enum.String()` as the error message.
    - `grpcStatus`: convert enum into grpc errors, `int32(enum)` as the code.
    - `number`: convert enum into `int32` and `int64`.
    - `validation`: verify whether the enum is in a valid range.
- `enums`: optional, a list of enums split by `+` for code generation, generate for all enums if not specified.
