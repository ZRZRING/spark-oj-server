# OVERVIEW.md

## go-judge API

- 编译请求

```json
{
    "cmd": [{
        "args": ["/usr/bin/g++", "a.cc", "-o", "a"],
        "env": ["PATH=/usr/bin:/bin"],
        "files": [{
            "content": ""
        }, {
            "name": "stdout",
            "max": 10240
        }, {
            "name": "stderr",
            "max": 10240
        }],
        "cpuLimit": 10000000000,
        "memoryLimit": 104857600,
        "procLimit": 50,
        "copyIn": {
            "a.cc": {
                "content": "#include <iostream>\nusing namespace std;\nint main() {\nint a, b;\ncin >> a >> b;\ncout << a + b << endl;\n}"
            }
        },
        "copyOut": ["stdout", "stderr"],
        "copyOutCached": ["a"]
    }]
}
```

```json
[
    {
        "status": "Accepted",
        "exitStatus": 0,
        "time": 303225231,
        "memory": 32243712,
        "runTime": 524177700,
        "files": {
            "stderr": "",
            "stdout": ""
        },
        "fileIds": {
            "a": "5LWIZAA45JHX4Y4Z" // 需要保存编译的二进制文件 id
        }
    }
]
```

- 运行请求

```json
{
    "cmd": [{
        "args": ["a"],
        "env": ["PATH=/usr/bin:/bin"],
        "files": [{
            "content": "1 1"
        }, {
            "name": "stdout",
            "max": 10240
        }, {
            "name": "stderr",
            "max": 10240
        }],
        "cpuLimit": 10000000000,
        "memoryLimit": 104857600,
        "procLimit": 50,
        "copyIn": {
            "a": {
                "fileId": "5LWIZAA45JHX4Y4Z" // 这个缓存文件的 ID 来自上一个请求返回的 fileIds
            }
        }
    }]
}
```

```json
[
    {
        "status": "Accepted",
        "exitStatus": 0,
        "time": 1173000,
        "memory": 10637312,
        "runTime": 1100200,
        "files": {
            "stderr": "",
            "stdout": "2\n"
        }
    }
]
```