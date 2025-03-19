# 第一阶段：构建应用
FROM golang:1.24.1-alpine3.21 AS builder

# 设置工作目录
WORKDIR /app

# 安装git和依赖
RUN apk add --no-cache git

# 复制go.mod和go.sum
COPY go.mod go.sum ./

# 复制项目文件
COPY . .

# 设置Go环境变量并构建应用
RUN go env -w GO111MODULE=on \
    && go env -w GOPROXY=https://goproxy.cn,direct \
    && go env -w CGO_ENABLED=0 \
    && go env -w GOFLAGS="-tags=prod" \
    && go env \
    && go mod tidy \
    && go build -o server .

# 第二阶段：运行阶段
FROM alpine:latest

# 安装基础工具和证书
RUN apk --no-cache add ca-certificates tzdata

# 设置时区
RUN cp /usr/share/zoneinfo/Asia/Shanghai /etc/localtime && \
    echo "Asia/Shanghai" > /etc/timezone

# 设置工作目录
WORKDIR /app

# 从构建阶段拷贝二进制文件
COPY --from=builder /app/server .

# 创建config目录，用于挂载配置文件
RUN mkdir -p /app/config

# 设置环境变量
ENV GIN_MODE=release
ENV CONFIG_PATH=/app/config/config.json

# 暴露端口
EXPOSE 8080

# 设置容器启动命令
# 使用软链接将配置文件链接到应用程序期望的位置
CMD ln -sf $CONFIG_PATH ./config.json && ./server 