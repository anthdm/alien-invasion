build:
	@go build -o bin/invasion ./src/...

run: build
	@./bin/invasion

test:
	@go test ./src/... --cover

clean:
	@rm -rf biN
