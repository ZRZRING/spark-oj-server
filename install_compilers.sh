#!/bin/bash

CONTAINER_NAME="go-judge"

echo "=========================================================="
echo "Installing compilers into docker container: $CONTAINER_NAME"
echo "=========================================================="

# 检查容器是否在运行
if ! docker ps | grep -q "$CONTAINER_NAME"; then
    echo "❌ 错误: 找不到正在运行的 '$CONTAINER_NAME' 容器。"
    echo "请先启动它: docker start $CONTAINER_NAME"
    exit 1
fi

echo "⏳ 正在检测容器内的系统类型并安装编译环境 (C/C++)..."

# 在容器内自动识别包管理器并执行安装任务
docker exec -u root "$CONTAINER_NAME" sh -c '
    if command -v apt-get >/dev/null 2>&1; then
        echo "👉 识别为 Debian/Ubuntu 系统，正在使用 apt-get 安装..."
        apt-get update -y
        apt-get install -y build-essential
    elif command -v apk >/dev/null 2>&1; then
        echo "👉 识别为 Alpine 系统，正在使用 apk 安装..."
        apk update
        apk add --no-cache build-base gcc g++ make
    else
        echo "❌ 无法识别包管理器 (未找到 apt-get 或 apk)。"
        exit 1
    fi
'

if [ $? -eq 0 ]; then
    echo "=========================================================="
    echo "✅ 安装成功！C/C++ 编译环境已就绪。"
    echo "🔍 g++ 版本信息："
    docker exec "$CONTAINER_NAME" g++ --version | head -n 1
else
    echo "=========================================================="
    echo "❌ 安装失败，请检查上方日志输出的错误信息。"
    exit 1
fi