clean:
	rm coverage.out

test:
	go test -cover ./...

coverage:
	go test -coverprofile=coverage.out ./...
	go tool cover -html=coverage.out

lint:
	golangci-lint run --enable-all --disable=gochecknoinits,gochecknoglobals,scopelint,gomnd,lll