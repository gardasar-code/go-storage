.PHONY: run

build:
	go build -o bin/main cmd/app/main.go

run:
	go run cmd/app/main.go