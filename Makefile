install-lint:
	go get github.com/golangci/golangci-lint/cmd/golangci-lint@v1.39.0

run-linter:
	sudo $(shell go env GOPATH)/bin/golangci-lint run

test:
	cd methods && go test -cover
