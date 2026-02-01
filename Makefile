# 宝塔面板MCP项目Makefile

# 设置变量
GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOMOD=$(GOCMD) mod
BINARY_NAME=mcp-btpanel
BUILD_DIR=build

# 系统信息
GOOS=$(shell go env GOOS)
GOARCH=$(shell go env GOARCH)

# 主文件路径
MAIN_PATH=main.go

# 默认目标
.PHONY: all
all: build

# 构建应用
.PHONY: build
build:
	@echo "编译项目..."
	@mkdir -p $(BUILD_DIR)
	@$(GOBUILD) -trimpath -ldflags "-s -w" -o $(BUILD_DIR)/$(BINARY_NAME) $(MAIN_PATH)
	@echo "编译完成: $(BUILD_DIR)/$(BINARY_NAME)"

# Windows特定构建
.PHONY: build-windows
build-windows:
	@echo "构建Windows版本..."
	@mkdir -p $(BUILD_DIR)
	@GOOS=windows GOARCH=amd64 $(GOBUILD) -trimpath -ldflags "-s -w" -o $(BUILD_DIR)/$(BINARY_NAME).exe $(MAIN_PATH)
	@echo "Windows版本构建完成: $(BUILD_DIR)/$(BINARY_NAME).exe"

# 清理构建产物
.PHONY: clean
clean:
	@echo "清理..."
	@rm -rf $(BUILD_DIR)
	@$(GOCLEAN)
	@echo "清理完成"

# 运行应用
.PHONY: run
run:
	@echo "运行应用..."
	@$(GOCMD) run $(MAIN_PATH)

# 运行测试
.PHONY: test
test:
	@echo "运行测试..."
	@$(GOTEST) -v ./...

# 整理依赖
.PHONY: tidy
tidy:
	@echo "整理依赖..."
	@$(GOMOD) tidy
	@echo "依赖整理完成"

# 帮助信息
.PHONY: help
help:
	@echo "可用的命令:"
	@echo "  make build          - 构建应用"
	@echo "  make build-windows  - 构建Windows版本"
	@echo "  make clean          - 清理构建产物"
	@echo "  make run            - 运行应用"
	@echo "  make test           - 运行测试"
	@echo "  make tidy           - 整理Go模块依赖"
	@echo "  make help           - 显示帮助信息" 