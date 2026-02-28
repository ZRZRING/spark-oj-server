# AGENT.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

请称呼我为 zrzring，总是用中文回复我

## 项目概述

Spark-OJ 是一个在线判题系统，用于编程竞赛。后端使用 GoFrame 框架，前端使用 Vue 3，集成 go-judge 进行代码执行。

## 架构

### 后端 (Go + GoFrame)
- **框架**: GoFrame v2.9.0
- **数据库**: PostgreSQL
- **缓存**: Redis
- **判题系统**: go-judge (运行在 Docker 的 5050 端口)
- **端口**: 8000

### 前端 (Vue 3)
- **框架**: Vue 3 + TypeScript + Vite
- **UI 库**: Element Plus
- **状态管理**: Pinia
- **代理**: 开发环境 API 请求代理到 `http://127.0.0.1:8000`

### 核心目录

**后端:**
- `api/v1/` - API 接口定义 (admin, contest, core, problem, submission, user)
- `internal/controller/v1/` - 控制器实现
- `internal/dao/` - 数据访问对象 (由 GoFrame CLI 自动生成)
- `internal/middleware/` - CORS、JWT 认证、权限控制
- `internal/model/entity/` - 数据库实体 (自动生成)
- `internal/service/` - 业务逻辑 (代码执行服务在 `execute_single_code.go`)
- `internal/router/` - 路由绑定
- `manifest/config/config.yaml` - 主配置文件
- `hack/config.yaml` - GoFrame CLI 配置，用于 DAO 生成

**前端:**
- `web/src/views/` - 页面组件
- `web/src/components/` - 可复用组件
- `web/src/stores/` - Pinia 状态管理
- `web/src/utils/service.ts` - Axios 实例，包含认证拦截器

## 常用开发命令

### 后端

```bash
# 运行开发服务器
gf run dev

# 从数据库生成 DAO
gf gen dao
```

### 前端

```bash
cd web

# 安装依赖
pnpm install

# 运行开发服务器
pnpm dev

# 生产构建
pnpm build

# 类型检查
pnpm type-check
```

### 判题系统 (go-judge)

```bash
# 首次启动 - 创建容器并安装编译器
docker run -it --privileged --shm-size=256m -p 5050:5050 --name=go-judge criyle/go-judge

# 启动已有容器
docker start -a go-judge
```

## API 设计模式

项目使用 GoFrame 约定：
- `api/v1/*.go` - 定义请求/响应结构体和路由元数据
- `internal/controller/v1/v1_*.go` - 实现处理器
- 路由器自动将控制器绑定到路径

核心端点：
- `POST /run` - 执行代码（交互式，用于题目编辑器）
- `POST /judge` - 提交代码判题（返回提交 ID，异步判题）
- `POST /upload_test_case` - 上传测试用例

## 代码执行服务

位于 `internal/service/execute_single_code.go`：
- 支持语言：cpp, c, java, python
- 编译型语言：先编译，缓存可执行文件，再运行
- 解释型语言：直接运行
- 与 go-judge REST API 通信，地址：`http://localhost:5050/run`

## 配置

- `manifest/config/config.yaml` - 数据库、Redis、判题服务 URL、第三方 API
- `web/.env.development` - 前端 API 基础 URL
- 判题 API 基础 URL 在 `config.yaml` 的 `thirdParty.judge.baseUrl` 配置

## 数据库实体

核心表：
- `problem` - 题目元数据（标题、限制、难度）
- `problem_content` - 题目描述（HTML 内容）
- `submission` - 代码提交及结果
- `contest` - 比赛信息
- `contest_problem` - 比赛-题目关联
- `user_base` - 用户信息
- `user_role` - 用户角色（管理员等）
- `tag` - 题目标签

## 认证

- 基于 JWT 的认证
- 中间件在 `internal/middleware/jwt_auth.go`
- 受保护的路由在 `/admin` 前缀下
- 前端在 localStorage 存储 token，请求时添加 `Authorization: Bearer <token>` 头