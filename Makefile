all:	install

clean:
	go clean -i ./...

install:
	go install .
