EXPORTER_NAME=cloudflare-exporter
VERSION=v1.0.0

build:
	go build -o bin/${EXPORTER_NAME}

run: build
	./bin/${EXPORTER_NAME}

clean:
	rm -rf ./bin
