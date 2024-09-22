FROM golang:1.23.0-alpine3.19
EXPOSE 5001
WORKDIR /app
RUN mkdir /build
COPY . /build/
RUN cd /build && rm go* && \
	go mod init cloudflare-status && go mod tidy && \
	apk add make && \
	make build && cp /build/bin/cloudflare-exporter /app/ && \
	apk del make && \
	rm -rf /build && rm -vrf /var/cache/apk/* rm -rf /go/pkg
ENTRYPOINT /app/cloudflare-exporter
