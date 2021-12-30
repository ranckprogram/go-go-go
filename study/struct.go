package main

import (
	"fmt"
)



type Box struct {
	Length int
	Width int
}


func (b *Box) getArea() int {
	return b.Length * b.Width
}


func main () {
	b := Box{6,2}
	fmt.Println(b.getArea())
}