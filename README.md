# go-go-go

### 环境安装

### 包管理工具

1. 安装指令

```
go get github.com/gin-gonic/gin
```

2. 构建exe

```
go build main.go
```

3. 运行

```
go run main.go
```


### wasm

```
GOOS=js GOARCH=wasm go build -o lib.wasm wasm.go
```

- [ ] go中定义方法注入浏览器 window对象 中调用
- [ ] go中调用浏览器中的方法
