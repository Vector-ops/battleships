build:
	@go build -o bin/bts -v

run: build
	@./bin/bts

test:
	@go test -v ./...

clean:
	rm -rf ./bin