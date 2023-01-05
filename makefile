BINARY_NAME = rocketfeed

all: clean build

build:
	GOOS=linux GOARCH=386 go build -o build/${BINARY_NAME}-linux-386
	GOOS=linux GOARCH=amd64 go build -o build/${BINARY_NAME}-linux-amd64
	GOOS=linux GOARCH=arm go build -o build/${BINARY_NAME}-linux-arm
	GOOS=linux GOARCH=arm64 go build -o build/${BINARY_NAME}-linux-arm64
	GOOS=darwin GOARCH=amd64 go build -o build/${BINARY_NAME}-darwin-amd64
	GOOS=darwin GOARCH=arm64 go build -o build/${BINARY_NAME}-darwin-arm64
	GOOS=windows GOARCH=386 go build -o build/${BINARY_NAME}-windows-386
	GOOS=windows GOARCH=amd64 go build -o build/${BINARY_NAME}-windows-amd64

clean:
	rm -f build/${BINARY_NAME}-*

.PHONY: build