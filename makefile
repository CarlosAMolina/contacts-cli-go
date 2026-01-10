test: fmt
	go test ./src/...

fmt:
	go fmt ./src/...

run:
	go run src/main.go src/types.go

tidy:
	go mod tidy
