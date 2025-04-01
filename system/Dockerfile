FROM ubuntu:latest
LABEL authors="bingodfok"
# 设置环境变量
ENV GOOS=linux \
    GOARCH=amd64 \
    CGO_ENABLED=0 \
    GOPROXY=https://goproxy.cn,direct
# 工作目录
WORKDIR /app

EXPOSE 9000

CMD ["./system"]