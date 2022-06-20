.PHONY: build
build:
	go build -o ./bin/dlr ./main.go

.PHONY: build-linux
build-linux:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o ./bin/dlr ./main.go