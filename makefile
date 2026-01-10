build:
	go build -o contacts-cli-go ./src

fmt:
	go fmt ./src/...

run:
	go run ./src/ src/fake.json 234

test: fmt
	go test ./src/...

tidy:
	go mod tidy

