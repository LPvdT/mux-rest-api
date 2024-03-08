DIR=dist
BIN=rest-api
OS=linux

.PHONY: build build-linux build-windows run clean

build: build-linux build-windows

build-linux: 
	GOARCH=amd64 GOOS=linux go build -o ${DIR}/${BIN}-linux

build-windows: 
	GOARCH=amd64 GOOS=windows go build -o ${DIR}/${BIN}-windows.exe

run: build
	./${DIR}/${BIN}-${OS}

clean:
	go clean
	rm dist/*
