BINARY_NAME=pokedexcli

build:
	@go build -o ${BINARY_NAME}

run:
	@go build -o ${BINARY_NAME}
	./${BINARY_NAME}

clean:
	@go clean
	@rm ${BINARY_NAME}

