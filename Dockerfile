FROM golang:1.16 as builder

ENV GO111MODULE=on \
    GOPROXY=https://goproxy.cn,direct

WORKDIR /ginchat
COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build  -ldflags="-w -s" -o main
RUN mkdir publish  \
    && cp main publish  \
    && cp -r config publish

FROM busybox:1.28.4

WORKDIR /ginchat

COPY --from=builder /ginchat/publish .

# 指定运行时环境变量
ENV GIN_MODE=release
EXPOSE 3000

ENTRYPOINT ["./main"]
