build:
	go build -o ./bin/lingscoin

run: build
	./bin/lingscoin

test:
	go test -v ./...