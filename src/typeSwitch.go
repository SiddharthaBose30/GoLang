package main

import (
	"fmt"
)

func type_switch(i interface{}) {
	switch i.(type) {
	case string:
		fmt.Printf("%T,%v\n", i, i)
	case int:
		fmt.Printf("%T,%v\n", i, i)
		x := i.(int)
		fmt.Println(x * x)
	default:
		fmt.Printf("Nothing")
	}

}
func main() {
	type_switch(12)
	type_switch("213")
	type_switch(true)
	// type x interface{}

}
