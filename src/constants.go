package main

import "fmt"

// const ms int32 = 123
const (
	a = iota
	b
	c
)
const (
	x = iota
	y
	z
)
const (
	_ = iota
	j
	k
)
const (
	q = 1 << (10 * iota)
	r
	s
)

func main() {
	// const ms int8 = 123
	// const str = "123"
	// const flt = 123.23
	// const bl = true
	// // fmt.Printf("%v %T\n", ms, ms)
	// fmt.Printf("%v %T\n", str, str)
	// fmt.Printf("%v %T\n", flt, flt)
	// fmt.Printf("%v %T\n", bl, bl)
	fmt.Printf("%v %T\n", a, a)
	fmt.Printf("%v %T\n", b, b)
	fmt.Printf("%v %T\n", c, c)
	// fmt.Printf("%v %T\n", x, x)
	// fmt.Printf("%v %T\n", j, j)
	fmt.Printf("%v %T\n", q, q)
	fmt.Printf("%v %T\n", r, r)
	fmt.Printf("%v %T\n", s, s)
}
