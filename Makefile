EXPORTER_NAME=cloudflare-exporter
CMD_DIR=cmd
VERSION=v1.2.1
BUILD_DIR=build

build:
	go build -C ${CMD_DIR}/${EXPORTER_NAME}  -o ../../${BUILD_DIR}/${EXPORTER_NAME}
	tar -czf ${BUILD_DIR}/${EXPORTER_NAME}-${VERSION}.tar.gz ${BUILD_DIR}/${EXPORTER_NAME}

run: build
	./${BUILD_DIR}/${EXPORTER_NAME}

clean:
	rm -rf ${BUILD_DIR}
