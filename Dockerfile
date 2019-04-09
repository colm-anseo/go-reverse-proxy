# latest golang version as of 2019/04/09 - see https://golang.org/dl/
FROM golang:1.12.3 AS go-build-image

WORKDIR /go/src/go-reverse-proxy
COPY *.go ./
RUN go get
RUN CGO_ENABLED=0 go build

##########
FROM scratch

COPY --from=go-build-image \
	go/src/go-reverse-proxy/go-reverse-proxy /app/

CMD ["/app/go-reverse-proxy"]
EXPOSE 9999
