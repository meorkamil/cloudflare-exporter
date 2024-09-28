EXPORTER_NAME=cloudflare-exporter
VERSION=v1.0.0
BUILD_DIR=build

build:
	go build -C cmd  -o ../${BUILD_DIR}/${EXPORTER_NAME}

run: build
	./${BUILD_DIR}/${EXPORTER_NAME}

clean:
	rm -rf ${BUILD_DIR}
