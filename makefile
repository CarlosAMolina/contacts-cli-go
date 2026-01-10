build:
	go build -o contacts-cli-go ./src

test: fmt
	go test ./src/...

fmt:
	go fmt ./src/...

run:
	go run ./src/ src/fake.json 234

tidy:
	go mod tidy
