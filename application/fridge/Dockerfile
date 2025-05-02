FROM golang:1.24.0 AS builder
LABEL authors="bingodfok"

WORKDIR /builds

ENV GOOS=linux \
    GOARCH=amd64 \
    CGO_ENABLED=0 \
    GOPROXY=https://goproxy.cn,direct

COPY . .

RUN go mod tidy
RUN go build -o freshguard-fridge .

FROM alpine:latest

WORKDIR /app

COPY --from=builder /builds/freshguard-fridge .
COPY --from=builder /builds/etc ./etc

EXPOSE 8081

CMD ["./freshguard-fridge"]