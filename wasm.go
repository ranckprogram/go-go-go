package main

import (
	"fmt"
	"syscall/js"
)

func foo( args []js.Value) {
	fmt.Println("hellow wasm")
	fmt.Println(args)
}


var cb = js.FuncOf(func(this js.Value, args []js.Value) interface{} {
    fmt.Println("button clicked")
    // cb.Release() // release the function if the button will not be clicked again
    return nil
})


func main() {
	fmt.Println("Hello, Go WebAssembly!")


	js.Global().Set("aaaload", cb )

  js.Global().Set("value", 100)
	select {}

}
