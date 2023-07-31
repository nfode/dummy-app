gen:
	go generate ./...

build: gen
	go build -o bin/dummy_app ./cmd/dummy_app
run: build
	./bin/dummy_app


