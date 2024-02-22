.DEFAULT_GOAL := run

fmt:
		go fmt

check: fmt
		goimports -l -w .

run: check
		go run hello.go