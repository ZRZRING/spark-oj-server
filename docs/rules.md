# 代码规范

- 保证 err 变量为框架原生操作的返回或者 `gerror.NewCode(code, [msg])`，并记录日志 `g.Log().Error(ctx, err, [...])`
- db/lduoj.sql 是其他项目的 DDL，无参考价值，请直接在 internal/model 查看数据库模型定义