# 项目开发进度

# 项目开发帮助

项目启动

```bash
gf run dev
```

评测机使用开源项目 go-judge

```bash
# 首次启动创建容器，需要进去安装 g++ 等编译工具
docker run -it --privileged --shm-size=256m -p 5050:5050 --name=go-judge criyle/go-judge
# 后续启动容器
docker start -a go-judge
```