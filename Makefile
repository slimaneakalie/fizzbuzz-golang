build:
	GO111MODULE=on go build \
     -v -o ./service cmd/main/*.go

clean:
	rm ./service || true

run: build
	./service

doc:
	@which swagger || GO111MODULE=on go get -u -v github.com/go-swagger/go-swagger@latest
	swagger validate api/swagger.yml
	swagger serve api/swagger.yml

testing:
	go test -v ./...