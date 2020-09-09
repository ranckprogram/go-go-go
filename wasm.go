package main

import (
	"fmt"
	"syscall"
)


func main() {
	fmt.Println("Hello, Go WebAssembly!")
	registerCallbacks()
	select {}


}

func registerCallbacks() {
	syscall.js.Global().Set("agetOne", syscall.js.FuncOf(agetOne))
	syscall.js.Global().Set("getUserName", syscall.js.FuncOf(getUserName))
}

func agetOne(this syscall.js.Value, args []syscall.js.Value) interface{} {
	return 2
}

func getUserName (this syscall.js.Value, args []syscall.js.Value) interface{} {
	 username := getSystemUserName()
	return username
}

func getSystemUserName() string {
	var size uint32 = 128
	var buffer = make([]uint16, size)
	user := syscall.StringToUTF16Ptr("USERNAME")
	domain := syscall.StringToUTF16Ptr("USERDOMAIN")
	r, err := syscall.GetEnvironmentVariable(user, &buffer[0], size)
	if err != nil {
			return ""
	}
	buffer[r] = '@'
	old := r + 1
	if old >= size {
			return syscall.UTF16ToString(buffer[:r])
	}
	r, err = syscall.GetEnvironmentVariable(domain, &buffer[old], size-old)
	return syscall.UTF16ToString(buffer[:old+r])
}