build:
	CGO_ENABLED=0 GOOS=linux go build -ldflags "-X 'github.com/flagship-io/flagship/cmd/version.Version=${FLAGSHIP_VERSION}'" -o flagship

test: SHELL:=/bin/bash
test:
	mkdir -p coverage
	go test -v -race `go list ./... | grep -v cmd/analyze/code-samples | grep -v cmd/configuration` -coverprofile coverage/cover.out.tmp
	cat coverage/cover.out.tmp | grep -v "mock_\|cmd/analyze/code-samples" > coverage/cover.out
	go tool cover -html=coverage/cover.out -o coverage/cover.html
	go tool cover -func=coverage/cover.out