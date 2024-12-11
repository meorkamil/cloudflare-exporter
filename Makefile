EXPORTER_NAME=cloudflare-exporter
CMD_DIR=cmd
VERSION=v1.2.3
BUILD_DIR=build
CONFIG_PATH=./config/config.yml

debug:
	cd ${CMD_DIR}/${EXPORTER_NAME} && go run main.go
build:
	go build -C ${CMD_DIR}/${EXPORTER_NAME} -C ${CMD_DIR}/${BIN_NAME} -o ../../${BUILD_DIR}/${EXPORTER_NAME}
	tar -czf ${BUILD_DIR}/${EXPORTER_NAME}-${VERSION}.tar.gz ${BUILD_DIR}/${EXPORTER_NAME}

run: build
	./${BUILD_DIR}/${EXPORTER_NAME}

clean:
	rm -rf ${BUILD_DIR}
