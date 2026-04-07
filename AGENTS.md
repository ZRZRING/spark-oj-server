# AGENTS.md

本文件为 AI 编程助手提供在 SparkOJ 仓库中工作的指导。

## 项目概述

SparkOJ 是一个在线评测系统（Online Judge），本项目为其后端。

- **框架**: Go (GoFrame v2.10) + PostgreSQL 18.3 + Redis 8.6.1
- **评测引擎**: go-judge（Docker 容器，HTTP API 通信）
- **入口**: `main.go` → `cmd.Main` → `internal/router/router.go`
- **端口**: 8000，公开路由 `/`，管理路由 `/admin`（JWT 鉴权）

## 关键目录

| 路径 | 说明 |
|---|---|
| `api/v1/` | API 请求/响应结构定义（`g.Meta` 标签定义路由） |
| `internal/controller/v1/` | 请求处理器（由 `make ctrl` 生成骨架） |
| `internal/dao/` | 数据访问对象（`gf gen dao` 自动生成，`internal/` 子目录禁止编辑） |
| `internal/model/` | Entity/DO 结构（自动生成，禁止编辑） |
| `internal/middleware/` | JWT 鉴权、CORS、日志等中间件 |
| `internal/router/` | 路由注册 |
| `pkg/enums/` | 枚举类型（JudgeStatus, Language, JudgeType 等） |
| `config/config.yaml` | 运行时配置 |
| `db/spark-oj.sql` | 数据库 DDL |
| `output/` | 日志输出目录 |

## 常用命令

```bash
go run main.go          # 启动服务
make build              # 构建二进制
make dao                # 从数据库生成 DAO/Entity/DO
make ctrl               # 从 API 定义生成 controller 骨架
make enums              # 生成枚举代码
```

## 开发流程

如果你需要开发一个新接口，你需要先在 `api/v1/` 目录下定义 API 定义，然后运行 `make ctrl` 生成 controller 骨架。然后，你需要在 `internal/controller/v1/` 目录下实现该接口的处理逻辑。

API 定义需要遵守 goframe-v2 规则，你可以查看 goframe-v2 这个 skills 获取帮助。

基本的 API 设计需要：

```go
type xxxReq struct {
	g.Meta `path:"/xxx" method:"POST" tags:"xxx" summary:"xxx"` 
	// 定义请求参数
}

type xxxRes struct {
	// 定义响应参数
}

// 如果有数组，需要定义 item 结构
type xxxItem struct {
	// 定义数组元素的结构
}
```

如果你需要修改数据库结构，你需要先在 `db/spark-oj.sql` 中定义 DDL，然后备份数据库并重新建表，然后删除 `internal/dao/` 和 `internal/model/entity/` 和 `internal/model/do/` 目录下的所有文件，最后运行 `make dao` 生成 DAO/Entity/DO。

## 重要规则

@docs/RULES.md
