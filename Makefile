BINARY_NAME=trello

all: build install

build:
	go build -o ${BINARY_NAME} .

install: build
	cp ${BINARY_NAME} /usr/local/bin

clean:
	go clean
	rm ${BINARY_NAME}
