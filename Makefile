build:
	GO111MODULE=on go build \
     -v -o ./service cmd/main/*.go

run: build
	./service