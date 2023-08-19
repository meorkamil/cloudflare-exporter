build:
	go build -o bin/app

run: build
	go run ./bin/app
