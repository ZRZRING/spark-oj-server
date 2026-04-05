# RULES.md

**如果你犯了错误被我指正，请总结并写入 `docs/RULES.md` 中。**

- 请称呼我为 zrzring，你是一个精通编程的AGENT，总是用中文回复我。
- 文档（尤其是CLAUDE.md）和代码不同时，以代码为准，并更新文档。
- 保证 err 变量为框架原生操作的返回或者 `gerror.NewCode(code, [msg])` 或 `gerror.Wrap(err, [msg])`，并在必要的地方记录日志 `g.Log().Error(ctx, err, [...])`
- db/lduoj.sql 是其他项目的 DDL，无参考价值，请直接在 internal/model 查看数据库模型定义
- 在 Windows 环境中出现 `apply_patch` rejected 时，优先判断为补丁上下文/换行问题，而不是用户拒绝修改；需先校验文件状态，再改用相对路径补丁或直接写文件完成变更。
- 设计 api 的时候，关于 id 的字段都设置为 string，即便数据库里存的是 number 等类型，这样空字符串好判断为空。
- 数据库里的 delete_at 框架会自动识别，不需要自己写代码实现 is null 判断。
- 备份数据库等操作可以放到 db 目录下。
- 你可以查看 /output 目录的日期最近的最新日志分析我最近遇到的问题，这样我就不用每次都发你 log 了，一般每次都加载最末尾的 50 行信息即可
- 文档默认保存到 /docs 目录下，写文档时严格按用户给的条目数量和结构来，不要自作主张拆分更多子条目。