build:
	GO111MODULE=on go build \
     -v -o ./service cmd/main/*.go

clean:
	rm ./service || true

run: build
	./service

testing:
	go test -v ./...