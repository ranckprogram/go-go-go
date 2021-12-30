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


go build -buildmode=c-shared -o go-ping.dll go-ping.go


### wasm

#### 获取 wasm_exec

就在go安装的目录下面，可以通过下面指令直接copy到目标文件夹

```
cp "$(go env GOROOT)/misc/wasm/wasm_exec.js" .
```

#### html中调用

```html
  <script src="./wasm_exec.js"></script>
  <script>
    const go = new Go();

    WebAssembly.instantiateStreaming(fetch("main.wasm"), go.importObject).then((result) => {
      const b = go.run(result.instance);
      console.log(result)
      console.log(b)
      b.then(res => {
        console.log(res)
      })
    });

```

#### 编写

main.go

```go
package main

import (
  "syscall/js"
)


func sum(this js.Value, args []js.Value) interface{} {
	var sum int
	for _, val := range args {
		sum += val.Int()
	}
	return sum
}



func registerCallbacks() {
	js.Global().Set("Asum", js.FuncOf(sum))
  // more 
}

func main() {
	fmt.Println("Hello, Go WebAssembly!")

	registerCallbacks()
	select {}

}

```

#### 编译wasm

```
GOOS=js GOARCH=wasm go build -o lib.wasm wasm.go

// path 可以更精确，减少copy
GOOS=js GOARCH=wasm go build -o ./static/main.wasm main.go
```

- [ ] go中定义方法注入浏览器 window对象 中调用
- [ ] go中调用浏览器中的方法




