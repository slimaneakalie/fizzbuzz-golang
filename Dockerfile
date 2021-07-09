FROM golang:1.16-alpine

WORKDIR /app

COPY go.mod .
COPY go.sum .
RUN go mod download

COPY . .
RUN GO111MODULE=on go build -v -o ./service cmd/main/*.go

EXPOSE 9000

CMD ["./service"]