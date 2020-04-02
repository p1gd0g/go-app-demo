build:
	@GOARCH=wasm GOOS=js go build -o app.wasm ./app
	@go build -o oschina ./server

run: build
	./oschina