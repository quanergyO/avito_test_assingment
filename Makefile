.PHONY: run, build, test, cover

run:
	docker-compose up avito-app

build:
	go build cmd/main.go

test:
	go test -v -count=1 ./...

cover:
	go test -short -count=1 -race -coverprofile=coverage.out ./...
	go tool cover -html=coverage.out
	rm coverage.out

gen:
	mockgen -source=internal/service/service.go -destination=internal/service/mocks/mock.go