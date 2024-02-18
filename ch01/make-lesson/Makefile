.DEFAULT_GOAL := run

fmt:
		go fmt

check: fmt
		goimports -l -w .

run: check
		go mod tidy
		go build
