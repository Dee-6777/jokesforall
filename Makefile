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
	$(BINARY)

test:
	go test ./cmd -v

joke:
	$(BINARY) 

choose:
	go build 
	@echo "Enter number of jokes you want to have:"; \
    read num; \
	$(BINARY) --num $$num

mod:
	go mod tidy 

clean:
	rm ${BINARY}
