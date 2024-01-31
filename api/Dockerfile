FROM golang:1.18-alpine as go-build

WORKDIR /app

COPY . .

RUN apk add -u -t build-tools curl git && \
    go build -o main main.go && \
    apk del build-tools && \
    rm -rf /var/cache/apk/*

FROM alpine:latest

WORKDIR /app

COPY --from=go-build /app /app/

EXPOSE 8080

ENTRYPOINT [ "/app/main" ]