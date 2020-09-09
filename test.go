package main

import "fmt"

func max(x, y int) (int) {
	if ( x > y) {
		return x
	} else {
		return y
	}
}

func main() {
   a:= max(5, 10)
	 fmt.Println(a)
}