FROM golang:1.16-alpine AS builder
#ENV GOPROXY="https://proxy.golang.org"
ENV GOPROXY="https://goproxy.io,direct"
ENV GOPROXY="https://goproxy.cn,direct"
COPY . /build
WORKDIR /build
RUN GO111MODULE="on" GOPROXY=$GOPROXY CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags "-s -w" web-archiver/cli/web-archiver
RUN sed -i "s@https://dl-cdn.alpinelinux.org/@https://mirrors.huaweicloud.com/@g" /etc/apk/repositories
RUN apk update && apk add upx
RUN upx web-archiver

FROM alpine
RUN mkdir -p /app
COPY --from=builder /build/web-archiver /app/
WORKDIR /app
CMD ./web-archiver