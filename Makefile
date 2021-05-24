build:
	go build -o helmplate cmd/helmplate/main.go

install:
	go get ./...

run: build
	./helmplate create