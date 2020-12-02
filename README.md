# go-go-go

学习文档
https://golang.google.cn/pkg/

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

4. 检测

```
go vet main.go
```


### wasm

```
GOOS=js GOARCH=wasm go build -o lib.wasm wasm.go
```

- [ ] go中定义方法注入浏览器 window对象 中调用
- [ ] go中调用浏览器中的方法

GOOS=js GOARCH=wasm go build -o ping.wasm ping.go



go build -buildmode=c-shared -o go-ping.dll go-ping.go
