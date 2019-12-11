SOURCEDIR=.

GOCMD=go
GOBUILD=$(GOCMD) build
BINARY=golang-seed-api
BINARY_PATH=$(SOURCEDIR)/cmd/golang-seed-api/main.go

.DEFAULT_GOAL: $(BINARY)

all: clean build

build: 
	$(GOBUILD) -o ${BINARY} $(BINARY_PATH)

clean: 
	rm -f $(BINARY)
