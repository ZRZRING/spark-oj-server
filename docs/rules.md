# 代码规范

- 保证 err 变量为框架原生操作的返回或者 `gerror.NewCode(code, [msg])`，并记录日志 `g.Log().Error(ctx, err, [...])`
- db/lduoj.sql 是其他项目的 DDL，无参考价值，请直接在 internal/model 查看数据库模型定义
- 在 Windows 环境中出现 `apply_patch` rejected 时，优先判断为补丁上下文/换行问题，而不是用户拒绝修改；需先校验文件状态，再改用相对路径补丁或直接写文件完成变更。
- 设计 api 的时候，关于 id 的字段都设置为 string，即便数据库里存的是 number 等类型，这样空字符串好判断为空。