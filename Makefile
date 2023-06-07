

all: run

build:
	go build -o bin/main cmd/main.go

run:
	go run cmd/main.go

test:
	go test -v ./...

dep:
	go mod download

vet:
	go vet ./...

lint:
	go fmt ./...


clean:
	rm -rf bin