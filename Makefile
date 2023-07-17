.PHONY: pb

pb:
	@protoc \
		-I ./pb \
		--go_out ./pb \
		--go_opt paths=source_relative \
		./pb/setter/*.proto \
		./pb/validator/*.proto \
		./pb/errorer/*.proto \
		./pb/consts/*.proto \
		./pb/i18n/*.proto \
		./pb/enums/*.proto \
		./pb/getter/*.proto
