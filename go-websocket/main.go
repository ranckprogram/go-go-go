package main

import (
	"fmt"
	"github.com/gorilla/websocket"
	"net/url"
	"syscall/js"
)


type DataSet struct {
	name string
	age int
}


func (d *DataSet) Set(name string, age int)  {
	d.name = name
	d.age = age
}

func sum(this js.Value, args []js.Value) interface{} {
	var sum int
	for _, val := range args {
		sum += val.Int()
	}
	// println(sum)
	return sum
}

func obj(this js.Value, args []js.Value) interface{} {
	// client := &http.Client{}
	// url := "http://www.ruanyifeng.com/home.html"
	// request, err := http.NewRequest("GET", url, nil)

	u := url.URL{Scheme: "ws", Host: "10.0.0.66:28080", Path: "/apis/ws.linkingthing.com/v1/alarm"}
	var	w,_,_ =	websocket.DefaultDialer.Dial(u.String(), nil)

	w.ReadMessage()
	fmt.Println(w)
	return nil
}


func registerCallbacks() {
	js.Global().Set("Asum", js.FuncOf(sum))
	js.Global().Set("Aobj", js.FuncOf(obj))
}

func main() {
	fmt.Println("Hello, Go WebAssembly!")

	// js.Global().Set("AAAVObj", obj)

  js.Global().Set("AAAvalue", 100)
	registerCallbacks()

	select {}

}
