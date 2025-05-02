FROM golang:1.24.0 AS builder
LABEL authors="bingodfok"

ENV GOOS=linux \
    GOARCH=amd64 \
    CGO_ENABLED=0 \
    GOPROXY=https://goproxy.cn,direct

WORKDIR /builds

COPY . .

RUN go mod tidy
RUN go build -o freshguard_system .

FROM alpine:latest

WORKDIR /app

COPY --from=builder /builds/freshguard_system .
COPY --from=builder /builds/etc ./etc

EXPOSE 8081

CMD ["./freshguard_system"]