FROM golang:1.23.0-alpine3.19
EXPOSE 5001
WORKDIR /app
COPY . /app
RUN cd /app && \
	apk add make && \
	make build && \
	rm -vrf /var/cache/apk/* rm -rf /go/pkg
ENTRYPOINT /app/build/cloudflare-exporter
