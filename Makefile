.PHONY: build run test clean

build:
	go build -o api main.go

run:
	go run main.go

test:
	go test ./...

clean:
	rm -f api
	go clean
	rm -rf docs

swagger:
	swag init -g main.go

deps:
	go mod download
	go mod tidy

docker-build:
	docker-compose build

docker-up:
	docker-compose up -d

docker-down:
	docker-compose down

docker-logs:
	docker-compose logs -f api

docker-restart:
	docker-compose restart api

