FROM golang:1.24.0 AS builder
LABEL authors="bingodfok"

ENV GOOS=linux \
    GOARCH=amd64 \
    CGO_ENABLED=0 \
    GOPROXY=https://goproxy.cn,direct

WORKDIR /builds

COPY . .
COPY ./application/resources/etc ./etc

RUN go mod tidy
RUN go build -o freshguard_resources ./application/resources/resources.go

FROM alpine:latest

WORKDIR /app

COPY --from=builder /builds/freshguard_resources .
COPY --from=builder /builds/etc ./application/resources/etc

EXPOSE 8082

RUN apk add dumb-init
ENTRYPOINT ["/usr/bin/dumb-init", "--"]

CMD ["./freshguard_resources"]