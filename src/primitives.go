package main

import (
	"fmt"
)

func main() {
	s := "sad"
	b := []byte(s)
	fmt.Printf("%v %T\n", s, s[2])
	fmt.Println(string(s[2]))
	fmt.Println(b)
	var a rune = 'a'
	fmt.Printf("%v %T\n", a, a)
}
