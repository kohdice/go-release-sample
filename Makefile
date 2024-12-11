SRC := .
BIN := sample
BUILD_LDFLAGS := "-s -w"

.PHONY: build
build:
	go build -trimpath -ldflags=$(BUILD_LDFLAGS) -o $(BIN) $(SRC)

.PHONY: clean
clean:
	go clean

.PHONY: fmt
fmt:
	go fmt ./...

.PHONY: lint
lint:
	go vet ./...

.PHONY: tidy
tidy:
	go mod tidy
