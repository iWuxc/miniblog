#定义全局Makefile 变量方便后边使用

COMMON_SELF_DIR := $(dir $(lastword $(MAKEFILE_LIST)))
#项目根目录
PROJ_ROOT_DIR := $(abspath $(shell cd $(COMMON_SELF_DIR) && pwd -P))
#构建产物、临时文件存放目录
OUTPUT_DIR := $(PROJ_ROOT_DIR)/_output

#============================================================================================
# 定义默认目标
.DEFAULT_GOAL := all

#定义Makefile all伪目标，执行 make 时，默认执行该目标
.PHONY: all
all: tidy format build add-copyright

#定义Makefile其他目标，执行 make clean 时，执行该目标
.PHONY: build
build: tidy
	@go build -v -o $(OUTPUT_DIR)/mb-apiserver $(PROJ_ROOT_DIR)/cmd/mb-apiserver/main.go

.PHONY: format
format:
	@gofmt -s -w ./

.PHONY: add-copyright
add-copyright:
	addlicense -v -f $(PROJ_ROOT_DIR)/scripts/boilerplate/boilerplate.txt $(PROJ_ROOT_DIR) --skip-dir=third_party,vendor.$(OUTPUT_DIR)

.PHONY: tidy
tidy:
	@go mod tidy

.PHONY: clean
clean:
	@rm -vrf $(OUTPUT_DIR)