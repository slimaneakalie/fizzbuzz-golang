run:
	docker-compose up -d  --force-recreate --build

doc:
	@which swagger || GO111MODULE=on go get -u -v github.com/go-swagger/go-swagger@latest
	swagger validate api/swagger.yml
	swagger serve api/swagger.yml

testing:
	go test -v ./...