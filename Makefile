VERSION=0.0.1
BINARY_NAME=webapp-tool
BUILD_PATH=build
BINARY_PATH=$(BUILD_PATH)/$(BINARY_NAME)

.PHONY: build

build:
	go build -o ${abspath ${BINARY_PATH}} main.go

exec:
	cd build && ./webapp-tool ../config.json

clean:
	go clean
	rm ${BINARY_PATH}