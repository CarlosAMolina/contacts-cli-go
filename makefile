test: fmt
	go test ./src/...

fmt:
	go fmt ./src/...

run:
	go run ./src/ 234

tidy:
	go mod tidy
