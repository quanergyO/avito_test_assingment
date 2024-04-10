FROM golang:latest

RUN go version
ENV GOPATH=/

COPY ./ ./
RUN go mod download
RUN go build -o avito-app ./cmd/main.go

CMD ["./avito-app"]
