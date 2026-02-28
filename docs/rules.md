# 代码规范

### 错误处理

- 保证 err 变量为框架原生操作的返回或者 `gerror.NewCode(code, [msg])`，并记录日志 `g.Log().Error(ctx, err, [...])`