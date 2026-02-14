build:
	go build -o bin/bankplay ./cmd/api

run: build
	./bin/bankplay

test:
	go test ./...

vet:
	go vet ./...

tidy:
	go mod tidy

clean:
	rm -rf bin

.PHONY: build run test vet tidy clean