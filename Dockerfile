FROM golang:latest

# 设置工作空间
WORKDIR /go/src/app
#　拷贝项目
COPY . .
# 启用mod
ENV GO111MODULE on
# 设置代理
ENV GOPROXY https://goproxy.io
# 下载依赖
RUN go mod tidy -v
# 编译
RUN go build -o app .
# 端口
EXPOSE 9000
# 启动
ENTRYPOINT ["./app"]
