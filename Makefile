build:
	go build -o helmplate cmd/helmplate/main.go

build-install:
	go build -o /usr/local/bin/helmplate cmd/helmplate/main.go

install:
	go get ./...

run: build
	./helmplate create ingress --chart chart