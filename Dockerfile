FROM golang:1.15-alpine

WORKDIR /app

COPY go.mod .
COPY go.sum .
RUN go mod download

COPY . .
RUN GO111MODULE=on go build -v -o ./fizzbuzz-bin cmd/main/*.go

EXPOSE 9000

CMD ["./fizzbuzz-bin"]