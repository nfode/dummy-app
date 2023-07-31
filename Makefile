gen:
	go generate ./...

build: gen
	go build ./cmd/dummy_app
run: build
	./dummy_app


