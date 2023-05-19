# Protogen

Generate golang code from proto files.

Depends on code generated from `protoc-gen-go`.

## Usage

```bash
go install github.com/guader/protogen/cmd/protoc-gen-go-ext

protoc -I ./pb \
	--go_out ./pb \
	--go_opt paths=source_relative \
	--go-ext_out ./pb \
	--go-ext_opt paths=source_relative \
	--go-ext_opt setter=true \
	--go-ext_opt setter_suffix=mysetter \
	./pb/*.proto

# Output files:
# - *.mysetter.go
```

## Options

- setter: `bool`, default: `false`, generate setter.
- setter_suffix: `string`, default: `"setter"`, suffix of generated filename, `<FILENAME>.<SUFFIX>.go`.