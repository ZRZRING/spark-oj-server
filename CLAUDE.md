# CLAUDE.md

本文件为 Claude Code (claude.ai/code) 及其他 AI 编程助手提供在 SparkOJ 仓库中工作的指导。

## 项目概述

SparkOJ 是一个在线评测系统（Online Judge）：
- **后端**: Go (GoFrame v2) + PostgreSQL + Redis
- **前端**: Vue 3 + TypeScript + Element Plus + Vite + Pinia（位于 `web/` 子目录）
- **评测引擎**: go-judge（通过 HTTP API 通信，需 Docker 容器运行）
- Go module: `spark-oj`

## 仓库结构

```
spark-oj/
├── main.go                  # Go 入口点
├── config.yaml              # 运行时配置 (数据库、Redis、评测 API)
├── api/v1/                  # API 请求/响应定义 (GoFrame OpenAPI 风格)
│   └── v1/{admin,contest,core,problem,submission,user}/
├── internal/
│   ├── cmd/                 # CLI 命令定义
│   ├── controller/v1/       # 请求处理器 (每个端点一个文件)
│   ├── dao/                 # 数据访问对象 (gf gen dao 自动生成)
│   │   └── internal/        # 自动生成的 DAO 内部实现 - 禁止编辑
│   ├── middleware/          # HTTP 中间件 (CORS, JWT, 日志)
│   ├── model/
│   │   ├── entity/          # 数据库实体结构 (自动生成 - 禁止编辑)
│   │   └── do/              # 数据输出结构 (自动生成 - 禁止编辑)
│   ├── router/              # 路由注册
│   └── service/             # 业务逻辑 (评测 API、代码执行、第三方接口)
├── pkg/
│   ├── consts/              # 全局常量
│   ├── enums/               # 枚举类型 (JudgeStatus, Language, JudgeType 等)
│   └── utils/               # 工具函数 (配置、比较)
├── hack/                    # GoFrame CLI 构建配置和 makefile
├── test/                    # 手动集成测试 (Go 脚本、.http 文件)
├── db/                      # 数据库备份
├── web/                     # Vue 3 前端 (独立子项目，你在这个项目里不要修改此项目)
├── docs/                    # 项目文档 (RULES.md 等)
├── output/                  # 日志输出目录
├── Makefile                 # 构建目标 (委托给 hack/*.mk)
└── docker-compose.yml       # Redis + go-judge 服务
```

## 构建 & 运行命令

### 后端 (Go)

```bash
# 启动服务器 (监听 :8000)
go run main.go

# 构建二进制
make build

# 从数据库生成 DAO/Entity/DO
make dao

# 从 API 定义生成 controller stubs
make ctrl

# 安装/更新 GoFrame CLI
make cli
```

### 前端 (Vue - 在 web/ 目录)

```bash
pnpm install     # 安装依赖
pnpm dev         # 开发服务器 (代理 /api 到 localhost:8000)
pnpm build       # 生产构建
pnpm type-check  # 类型检查
```

### 测试

无正式测试框架，手动测试脚本在 `test/` 目录，个人项目对大型测试的需求不大，跑通流程即可。

## 代码风格指南

### Go (后端)

**框架约定 (GoFrame v2):**
- API 定义使用 `g.Meta` 标签进行路由、验证和 OpenAPI 文档
- Controller 实现自动生成的接口 (定义在 `api/v1/v1.go`)
- 使用 `g.Log()`, `g.DB()`, `g.Client()`, `g.Cfg()` 获取框架服务
- DAO/Entity/DO 文件自动生成 - **禁止编辑** `internal/dao/internal/`、`internal/model/entity/`、`internal/model/do/`
- 使用 `gerror.NewCode(code, [msg])` 或 `gerror.Wrap(err, [msg])` 创建错误
- 始终记录错误日志: `g.Log().Error(ctx, err, [...额外上下文...])`
- 使用 `gconv.Struct()` 进行请求到模型的转换
- 业务验证错误使用 `gerror.NewCode(gcode.CodeInvalidRequest, "消息")`

**导入顺序:**
1. 标准库 (`context`, `fmt`, `encoding/json` 等)
2. 第三方库 (`github.com/gogf/gf/v2/...`, `github.com/golang-jwt/...`)
3. 项目内部导入 (`spark-oj/internal/...`, `spark-oj/pkg/...`, `spark-oj/api/...`)

**命名规范:**
- 文件: `snake_case.go`
- 结构体/类型: `PascalCase`
- 变量: `camelCase`
- 常量: `PascalCase` 或 `UPPER_SNAKE_CASE`
- DAO 变量: `PascalCase` (如 `dao.UserBase`, `dao.Problem`)

**API 设计:**
- API 请求/响应中的 ID 字段应为 `string` 类型（即使数据库存储数字）
- 空字符串表示未提供，避免 nil 混淆
- `delete_at` 由 GoFrame ORM 自动处理，无需手动添加 `IS NULL` 过滤

**错误处理模式:**
```go
if err != nil {
    g.Log().Error(ctx, err)
    return nil, err
}
```

### TypeScript / Vue (前端)

**技术栈:** Vue 3 Composition API + `<script setup lang="ts">` + Pinia stores + Vue Router 4

**命名规范:**
- `.vue` 文件: `PascalCase` (如 `ProblemPageView.vue`)
- `.ts` 文件: `camelCase` (如 `usePagedList.ts`)
- CSS class: `kebab-case` (如 `main-container`)
- TypeScript 通过 `transformKeysToCamelCase` 自动处理 snake_case 到 camelCase 的转换

**组件模式:**
```vue
<script setup lang="ts">
// 导入、composables、响应式状态、方法
</script>

<template>
  <!-- 优先使用 Element Plus 组件 -->
</template>

<style scoped>
/* 最小化自定义 CSS；优先使用 Element Plus 组件 */
</style>
```

**Store 模式 (Pinia):**
- 使用 `defineStore('name', () => { ... })` 定义 store (Composition API 风格)
- 在 store 旁边定义请求/响应接口
- 使用 `service.get<Req, Res>(url, req)` / `service.post<Req, Res>(url, req)` 调用 API
- 响应包装器: `{ code: number; message: string; data: T }`

**状态管理:**
- Token 和 username 存储在 `localStorage` 并与 Pinia refs 同步
- 使用 computed 属性派生状态 (如 `isLoggedIn`, `isAdmin`)

**关键约定:**
- 尽可能使用 Element Plus 组件以减少自定义 CSS
- 从前端角度看后端 API 是**只读的** - 建议修改而非直接修改后端
- 使用 `@/` 路径别名导入 (在 vite.config.ts 配置)
- 所有 API 响应在 axios 拦截器中经过 `transformKeysToCamelCase`

## 数据模型 (internal/model/entity/)

- **UserBase**: 用户表 (username, password, nickname, role, solved, submitted, rating)
- **Problem**: 题目表 (problem_id, title, judge_type, time_limit, memory_limit, content, rating)
- **Contest**: 比赛表 (contest_id, title, password, description, start_time, end_time, problems, practice)
- **Submission**: 提交表 (submission_id, problem_id, contest_id, username, result, language, code)
- **Resource**: 资源表
- **UserRole**: 用户角色表

## 重要规则

@docs/RULES.md

## go-judge API

评测机通过 HTTP API 通信，支持编译缓存。详细结构定义见 `internal/service/go_judge_api.go`。

**编译请求示例:**
```json5
{
    "cmd": [{
        "args": ["/usr/bin/g++", "a.cc", "-o", "a"],
        "env": ["PATH=/usr/bin:/bin"],
        "files": [{ "content": "" }, { "name": "stdout", "max": 10240 }, { "name": "stderr", "max": 10240 }],
        "cpuLimit": 10000000000,
        "memoryLimit": 104857600,
        "procLimit": 50,
        "copyIn": { "a.cc": { "content": "源代码..." } },
        "copyOut": ["stdout", "stderr"],
        "copyOutCached": ["a"]
    }]
}
```

**编译响应示例:**
```json5
[{ "status": "Accepted", "exitStatus": 0, "time": 303225231, "memory": 32243712,
   "files": { "stderr": "", "stdout": "" },
   "fileIds": { "a": "5LWIZAA45JHX4Y4Z" }  // 保存编译二进制的缓存 ID
}]
```

**运行请求示例 (使用编译缓存):**
```json5
{
    "cmd": [{
        "args": ["a"],
        "env": ["PATH=/usr/bin:/bin"],
        "files": [{ "content": "1 1" }, { "name": "stdout", "max": 10240 }, { "name": "stderr", "max": 10240 }],
        "cpuLimit": 10000000000,
        "memoryLimit": 104857600,
        "procLimit": 50,
        "copyIn": { "a": { "fileId": "5LWIZAA45JHX4Y4Z" } }  // 使用编译缓存的 fileId
    }]
}
```

**支持的编程语言:** C, C++, Java, Python

**判题状态:** Accepted, Wrong Answer, Time Limit Exceeded, Memory Limit Exceeded, Runtime Error, Compile Error

**判题类型:** Standard, Special Judge, Interactive, Communication, Output Only