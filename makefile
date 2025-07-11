# 设置编译目标
DIST := dist/server

# 编译Go程序
build:
	@echo "Building Go binary..."
	@CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o $(DIST) main.go

# 构建Docker镜像
docker-build:
	@echo "Building Docker image..."
	@docker build -f Dockerfile -t public:latest .

# 运行Docker容器
docker-run:
	@echo "Running Docker container..."
	@docker rm -f public
	@docker run -d --restart=always --net=host -v /root/cert:/app/cert --name public public

# 清理编译生成的二进制文件
clean:
	@echo "Cleaning up..."
	@rm -f $(DIST)

# 清理Docker镜像
docker-clean:
	@echo "Cleaning Docker images..."
	@docker rmi -f public:latest

# 定义一个目标来执行所有操作
all: build docker-build docker-run

dp:
	@CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o dist/headless headless/main.go
	@docker build -f Dockerfile.headless -t headless:latest .
	@docker rm -f headless
	@docker run -d --restart=always -p 8090:8090/tcp --name headless --init headless:latest

# 设置默认目标
.PHONY: build docker-build docker-run docker-clean clean all
.DEFAULT_GOAL := build