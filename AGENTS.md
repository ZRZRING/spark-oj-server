# AGENTS.md

本文件为 AI 编程助手提供在 SparkOJ 仓库中工作的指导。

## 项目概述

SparkOJ 是一个在线评测系统（Online Judge）：
- **后端**: Go (GoFrame v2) + PostgreSQL + Redis
- **前端**: Vue 3 + TypeScript + Element Plus + Vite + Pinia（位于 `web/` 子目录，是子仓库）
- **评测引擎**: go-judge（通过 HTTP API 通信，需 Docker 容器运行）
- Go module: `spark-oj`

## 构建 & 运行

```bash
go run main.go          # 启动服务器 (监听 :8000)
make build              # 构建二进制
make dao                # 从数据库生成 DAO/Entity/DO
make ctrl               # 从 API 定义生成 controller stubs
make enums              # 生成枚举代码
make service            # 生成 Service 代码
```

## 重要规则

@docs/RULES.md