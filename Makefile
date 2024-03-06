DIR=dist
BIN=rest-api
OS=linux

.PHONY: build run clean

build: build-linux build-windows

build-linux: ${DIR}/${BIN}-linux
	GOARCH=amd64 GOOS=linux go build -o ${DIR}/${BIN}-linux

build-windows: ${DIR}/${BIN}-windows.exe
	GOARCH=amd64 GOOS=windows go build -o ${DIR}/${BIN}-windows.exe

run: build
	./${DIR}/${BIN}-${OS}

clean:
	go clean
	rm dist/*
