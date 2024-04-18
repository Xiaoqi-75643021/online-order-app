# 使用官方Go镜像作为构建环境
FROM golang:1.22.1 AS builder

# 设置国内的Golang代理
ENV GOPROXY=https://goproxy.cn,direct

# 设置mysql数据库的环境变量
ENV ORDER_DB_HOST=127.0.0.1
ENV ORDER_DB_NAME=online-ordering-app
ENV ORDER_DB_PASSWORD=WXP2xiong!
ENV ORDER_DB_PROT=3306
ENV ORDER_DB_USER=root

# 工作目录
WORKDIR /app

# 复制go.mod和go.sum文件
COPY go.mod go.sum ./

# 下载依赖项
RUN go mod download

# 复制项目文件
COPY . .

# 设置工作目录为cmd目录(因为cmd目录包含了程序主入口main.go)
WORKDIR /app/cmd

# 构建应用程序
RUN CGO_ENABLED=0 GOOS=linux go build -o online-ordering-app

# 使用scratch作为最终镜像
FROM scratch

# 从builder镜像中复制构建的可执行文件到scratch镜像中
COPY --from=builder /app/cmd/online-ordering-app .

# 运行应用程序
CMD ["./online-ordering-app"]