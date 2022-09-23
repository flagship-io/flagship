build:
	CGO_ENABLED=0 GOOS=linux go build -o flagship

test: SHELL:=/bin/bash
test:
	mkdir -p coverage
	go test -v -race ./... -coverprofile coverage/cover.out.tmp -coverpkg=./...
	cat coverage/cover.out.tmp | grep -v "mock_\|examples" > coverage/cover.out
	go tool cover -html=coverage/cover.out -o coverage/cover.html
	go tool cover -func=coverage/cover.out