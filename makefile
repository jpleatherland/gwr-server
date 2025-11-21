.PHONY: run build

run:
	go run ./cmd/gwr-server

build:
	go build -o gwr-server ./cmd/gwr-server
