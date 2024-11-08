GO_FILE = main.go 
BINARY_NAME = main

all: build run

run: 
	go run $(GO_FILE)

build:
	go build -o $(BINARY_NAME) $(GO_FILE)

clean:
	rm -f $(BINARY_NAME)

.PHONY: all run build clean