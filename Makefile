.PHONY: run, build, test, cover

run:
	docker-compose up avito-app

build:
	go build cmd/main.go

test:
	docker run --name=redis-avito-test -p 6380:6379 -d --rm redis
	docker run --name=avito-test-bd -e POSTGRES_PASSWORD='qwerty' -p 5438:5432 -d --rm postgres
	go test -count=1 ./tests/ || true
	docker stop redis-avito-test
	docker stop avito-test-bd

cover:
	go test -short -count=1 -race -coverprofile=coverage.out ./...
	go tool cover -html=coverage.out
	rm coverage.out

gen:
	mockgen -source=internal/service/service.go -destination=internal/service/mocks/mock.go
