# 阶段 1：构建阶段
FROM golang:1.22-alpine AS builder

ENV CGO_ENABLED=0
ENV GOPATH=/go
ENV GOCACHE=/go-build

# 设置工作目录
WORKDIR /app

# 复制 go.mod 和 go.sum 文件
COPY go.mod go.sum ./

# 下载依赖
RUN --mount=type=cache,target=/go/pkg/mod/cache \
    GOPROXY=https://goproxy.cn,direct \
    go mod download

# 复制项目代码
COPY . .

# 编译 product api 服务
RUN go build -o product-api ./apps/product/api/product.go

# 编译 product rpc 服务
RUN go build -o product-rpc ./apps/product/rpc/product.go

# 编译 user api 服务
RUN go build -o user-api ./apps/user/api/user.go

# 编译 user rpc 服务
RUN go build -o user-rpc ./apps/user/rpc/user.go

# 阶段 2：运行时阶段
FROM alpine:latest

# 设置工作目录
WORKDIR /app

# 从构建阶段复制编译好的二进制文件
COPY --from=builder /app/product-api .
COPY --from=builder /app/product-rpc .
COPY --from=builder /app/user-api .
COPY --from=builder /app/user-rpc .


# 暴露端口
EXPOSE 8080 8081 8082 8083

