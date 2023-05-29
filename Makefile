.PHONY: pb

pb:
	@protoc \
		-I ./pb \
		--go_out ./pb \
		--go_opt paths=source_relative \
		./pb/setter/*.proto \
		./pb/validator/*.proto \
		./pb/errorer/*.proto
