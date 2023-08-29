# Code Compile
FROM golang:alpine AS build

# ENV GOPROXY=https://goproxy.cn
WORKDIR /app
COPY . /app

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 \
    go build -ldflags="-w -s" -o alert-webhook-receiver

# Image Build
FROM alpine:latest

ENV TZ=Asia/Dubai
# RUN echo "http://mirrors.aliyun.com/alpine/latest-stable/main/" > /etc/apk/repositories
RUN apk update --no-cache && \
    apk add --no-cache --update tzdata ca-certificates && \
    mkdir -p /app && \
    addgroup -S app && \
    adduser -S -u 1000 -g app -h /app app && \
    chown -R app:app /app && \
    cp /usr/share/zoneinfo/${TZ} /etc/localtime && \
    echo ${TZ} > /etc/timezone

COPY --from=build /app/alert-webhook-receiver /app/alert-webhook-receiver
RUN chmod +x /app/alert-webhook-receiver

USER 1000
WORKDIR /app
EXPOSE 9000

ENTRYPOINT ["/app/alert-webhook-receiver"]