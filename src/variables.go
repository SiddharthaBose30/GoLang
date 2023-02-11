package main

import (
	"fmt"
	"strconv"
)

var l int
var p float32 = 123.0

var (
	a = "Messi"
	b = "Sid"
	c = 123
)
var (
	f string = "asd"
)

func main() {
	// var i int
	// i = 12
	// var j int = 123
	// k := 12.0
	// fmt.Println("Hello World")
	// fmt.Println(i, j, k)
	// fmt.Printf("%v %T %T", i, j, k)\
	i := 123
	j := strconv.Itoa(i)
	// fmt.Printf("%T %T", a, b)
	// var a int = 123
	// var b = float32(b)
	fmt.Printf("%T %T", i, j)
	// fmt.Println()
	fmt.Printf("%T", c)

}
