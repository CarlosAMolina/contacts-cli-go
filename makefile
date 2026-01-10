test: fmt
	go test ./src/...

fmt:
	go fmt ./src/...

run:
	go run src/types.go src/search.go src/main.go 234

tidy:
	go mod tidy
