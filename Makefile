export GO111MODULE=on
export CGO_ENABLED=0

BINARY=jokesforall

all: build

build:
	go build 

install:
	go install

run:
	go build 
	./$(BINARY)

test:
	go test ./cmd -v

joke:
	go build 
	./$(BINARY) joke

mod:
	go mod tidy 

clean:
	rm ${BINARY}