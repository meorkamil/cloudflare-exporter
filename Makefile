build:
	go build -o bin/app

run: build
	./bin/app

clean:
	rm -rf ./bin
