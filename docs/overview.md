# Spark OJ

## 1. 项目概述

**Spark-OJ** 是一个在线评测系统（Online Judge），支持代码提交、自动评测、竞赛管理等功能。

### 1.1 技术栈

| 层级 | 技术选型 |
|------|----------|
| 后端框架 | GoFrame v2.9.0 (Go 1.24.1) |
| 数据库 | PostgreSQL |
| 缓存 | Redis |
| 前端框架 | Vue 3 + TypeScript |
| UI组件库 | Element Plus |
| 状态管理 | Pinia |
| 构建工具 | Vite |
| 评测引擎 | go-judge (Docker容器) |
| 认证 | JWT |

### 1.2 项目结构

```
spark-oj/
├── api/                    # API 定义层
│   └── v1/                 # v1 版本 API
│       ├── admin/          # 管理员接口
│       ├── contest/        # 竞赛接口
│       ├── core/           # 核心功能(评测/运行)
│       ├── problem/        # 题目接口
│       ├── submission/     # 提交记录接口
│       └── user/           # 用户接口
├── app/                    # 应用模块
│   └── go-judge/           # go-judge 客户端封装
│       ├── envexec/        # 执行环境
│       ├── env/            # 环境配置
│       ├── pb/             # protobuf 定义
│       └── worker/         # 工作池
├── db/                     # 数据库脚本
│   └── lduoj.sql           # 数据库结构
├── doc/                    # 文档
├── internal/               # 内部实现
│   ├── cmd/                # 命令入口
│   ├── controller/v1/      # 控制器实现
│   ├── dao/                # 数据访问层
│   ├── middleware/         # 中间件
│   ├── router/             # 路由配置
│   └── service/            # 业务服务层
├── manifest/               # 配置清单
│   ├── config/             # 配置文件
│   ├── docker/             # Docker配置
│   └── i18n/               # 国际化
├── resource/               # 静态资源
│   ├── public/             # 公共资源
│   └── static/             # 静态文件(测试用例等)
├── web/                    # 前端项目
│   └── src/
│       ├── components/     # 组件
│       ├── config/         # 配置
│       ├── router/         # 路由
│       ├── stores/         # Pinia状态管理
│       ├── utils/          # 工具函数
│       └── views/          # 页面视图
├── main.go                 # 程序入口
├── go.mod                  # Go依赖管理
└── Makefile                # 构建脚本
```

---

## 2. 后端架构

### 2.1 核心模块

#### 2.1.1 用户模块 (User)
- **登录认证**: `/login` POST
- **用户注册**: `/register` POST
- **个人资料**: `/profile/{username}` GET
- **训练统计**: `/training` GET

#### 2.1.2 题目模块 (Problem)
- **题目列表**: `/problem/list` GET
- **获取题目**: `/problem/{pid}` GET
- **创建题目**: `/problem/create` POST (需管理员权限)
- **更新题目**: `/problem/update` PUT (需管理员权限)

#### 2.1.3 提交模块 (Submission)
- **提交列表**: `/submission/list` GET
- **获取提交**: `/submission/{sid}` GET

#### 2.1.4 竞赛模块 (Contest)
- **竞赛列表**: `/contest/list` GET
- **获取竞赛**: `/contest/{cid}` GET
- **创建竞赛**: `/contest/create` POST
- **更新竞赛**: `/contest/update` PUT
- **竞赛排名**: `/contest/ranking` GET

#### 2.1.5 核心功能模块 (Core)
- **代码评测**: `/judge` POST
- **代码运行**: `/run` POST
- **上传测试用例**: `/upload/testcases` POST

### 2.2 中间件

| 中间件 | 功能 |
|--------|------|
| CORS | 跨域请求处理 |
| JWTAuth | JWT认证 |
| PermissionAuth | 权限验证 |

### 2.3 数据访问层 (DAO)

位于 `internal/dao/`:
- `user_base.go` - 用户基础信息
- `user_role.go` - 用户角色
- `problem.go` - 题目
- `contest.go` - 竞赛
- `contest_problem.go` - 竞赛题目关联
- `submission.go` - 提交记录
- `tag.go` - 标签
- `file.go` - 文件管理
- `judger.go` - 评测机

---

## 3. 评测系统

### 3.1 评测流程

```
用户提交代码 → 后端接收 → 评测队列 → go-judge容器 → 返回结果
```

### 3.2 支持的编程语言

| 语言 | 编译命令 | 运行命令 |
|------|----------|----------|
| C++ | g++ main.cpp -o main | ./main |
| C | gcc main.c -o main | ./main |
| Java | javac Main.java | java Main |
| Python | - | python3 main.py |

### 3.3 资源限制 (默认值)

| 资源 | 限制 |
|------|------|
| CPU时间 | 10秒 |
| 内存 | 256MB |
| 进程数 | 10 |
| 输出限制 | 10KB |

### 3.4 go-judge 集成

评测服务通过 HTTP API 与 go-judge 容器通信:
- 基础URL: `http://localhost:5050`
- 主要接口: `POST /run`

---

## 4. 数据库设计

### 4.1 核心数据表

| 表名 | 说明 |
|------|------|
| users | 用户信息 |
| problems | 题目 |
| contests | 竞赛 |
| contest_problems | 竞赛-题目关联 |
| contest_users | 竞赛-用户关联 |
| solutions | 提交记录 |
| tags / tag_pool | 标签系统 |
| roles / permissions | RBAC权限系统 |
| groups | 用户组 |

### 4.2 权限系统 (RBAC)

```
用户 → model_has_roles → 角色
角色 → role_has_permissions → 权限
用户 → model_has_permissions → 权限(直接)
```

---

## 5. 前端架构

### 5.1 页面路由

| 路由 | 页面 | 说明 |
|------|------|------|
| /home | HomeView | 首页 |
| /problems | ProblemsView | 题目列表 |
| /problem/:pid | ProblemView | 题目详情 |
| /contest | ContestsView | 竞赛列表 |
| /contest/:cid | ContestView | 竞赛详情 |
| /judge | SubmissionView | 提交记录 |
| /profile/:username | ProfileView | 用户资料 |
| /login | LoginView | 登录 |
| /register | RegisterView | 注册 |
| /admin/problems | ProblemsView | 管理题目 |
| /admin/problem/create | CreateProblemView | 创建题目 |

### 5.2 状态管理 (Pinia Stores)

| Store | 说明 |
|-------|------|
| user_store | 用户状态(登录/权限) |
| problem_store | 题目数据 |
| contest_store | 竞赛数据 |
| submission_store | 提交记录 |

### 5.3 API 请求

通过 `src/utils/service.ts` 封装 axios 请求。

---

## 6. 配置说明

### 6.1 后端配置 (`manifest/config/config.yaml`)

```yaml
server:
  address: ":8000"          # 服务端口
  openapiPath: "/api.json"  # OpenAPI文档
  swaggerPath: "/swagger"   # Swagger UI

database:
  default:
    link: "pgsql:用户:密码@tcp(主机:端口)/数据库"

redis:
  address: "127.0.0.1:6379"

thirdParty:
  judge:
    baseUrl: "http://localhost:5050"  # go-judge地址
```

### 6.2 前端配置

- 开发环境: `web/.env.development`
- 生产环境: `web/.env.production`

---

## 7. 启动方式

### 7.1 后端

```bash
# 开发模式
gf run dev

# 或直接运行
go run main.go
```

### 7.2 评测机

```bash
# 首次启动(需要安装g++等编译工具)
docker run -it --privileged --shm-size=256m -p 5050:5050 --name=go-judge criyle/go-judge

# 后续启动
docker start -a go-judge
```

### 7.3 前端

```bash
cd web
pnpm install
pnpm dev
```

---

## 8. 待完善功能

根据代码分析，以下功能尚未完全实现:

1. **Ranking功能**: 竞赛排名逻辑
2. **Codeforces集成**: 第三方平台数据同步
3. **部分管理功能**: 标签管理等

---
## 9. 开发建议

### 9.1 代码风格
- 使用 GoFrame 约定的目录结构
- API 定义在 `api/v1/` 目录
- 控制器实现在 `internal/controller/v1/`
- 业务逻辑在 `internal/service/`

### 9.2 添加新API流程
1. 在 `api/v1/` 定义请求/响应结构
2. 在 `api/v1/v1.go` 添加接口定义
3. 在 `internal/controller/v1/` 实现控制器
4. GoFrame 自动注册路由

### 9.3 前端开发
- 遵循 Vue 3 Composition API
- 使用 Element Plus 组件
- 状态管理使用 Pinia

---

## 10. 相关链接

- GoFrame 文档: https://goframe.org
- go-judge 项目: https://github.com/criyle/go-judge
- Vue 3 文档: https://vuejs.org
- Element Plus 文档: https://element-plus.org

