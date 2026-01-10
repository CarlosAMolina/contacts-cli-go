fmt:
	go fmt ./src/

run:
	go run src/main.go

test:
	go test ./src/

tidy:
	go mod tidy
