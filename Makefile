APP_NAME=aicom
SRC_DIR=./src
DIST_DIR=dist
GOPATH=$(shell go env GOPATH)

build:
	mkdir -p $(DIST_DIR)
	go build -o $(DIST_DIR)/$(APP_NAME) $(SRC_DIR)

install:
	make build
	mv $(DIST_DIR)/$(APP_NAME) $(GOPATH)/bin/

uninstall:
	rm $(GOPATH)/bin/$(APP_NAME)

test:
	go test ./src/...

clean:
	rm -rf $(DIST_DIR)
