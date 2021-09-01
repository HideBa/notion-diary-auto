lint:
	golangci-lint run --fix

test:
	go test -race -v ./..

build:
	go build ./main.go

run:
	air

db:
	docker-compose up -d db