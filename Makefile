EXPORTER_NAME=cloudflare-exporter
VERSION=v1.0.0

build:
	go build -o bin/${EXPORTER_NAME}

run: build
	./bin/${EXPORTER_NAME}
	#docker run --rm -it --entrypoint /app/${EXPORTER_NAME} -v "$(pwd)/bin:/app" golang:1.23.0-alpine3.19

clean:
	rm -rf ./bin
