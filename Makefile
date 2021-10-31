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

migrate:
	micreate create -ext sql -dir ./infrastructure/migrations -seq

migrate-apply:
	migrate -database "postgres://user:password@localhost:5432/dation-db?sslmode=disable" -path infrastructure/database/migrations up