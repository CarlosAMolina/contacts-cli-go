fmt:
	go fmt ./src/

test:
	go test ./src/...

tidy:
	go mod tidy
