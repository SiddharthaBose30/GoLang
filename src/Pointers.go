package main

import "fmt"

func main() {
	// a := 12
	// b := &a
	// var c *int = b
	// println(a, &b, *c)
	// var struct_pointer *app
	// struct_pointer = &app{Age: 12}
	// fmt.Println(*struct_pointer)
	// nil_and_new()
	// Arrays_Pointers()
	// Slices()
	Maps()
}

type app struct {
	Age int
}

func nil_and_new() {
	var ms *app
	fmt.Println(ms)
	ms = new(app)
	fmt.Println(*ms)
	(*ms).Age = 12
	fmt.Println(*&ms.Age == (*ms).Age && (*ms).Age == ms.Age)
	fmt.Println(ms.Age)
}

func Arrays_Pointers() {
	a := [3]int{1, 2, 3}
	b := &a[0]
	c := &a[1]
	d := a
	a[2] = 4
	fmt.Println(a, *b, *c, d)
}

func Slices() {
	//slices are nothing but a projection of arrays i.e they are pointers to an underlying array
	a := []int{1, 2, 4}
	a[1] = 2
	b := a
	a[1] = 123
	fmt.Println(a, b)
}

func Maps() {
	a := make(map[string]int)
	a = map[string]int{
		"Messi": 10,
	}
	b := a
	a["Messi"] = 20
	fmt.Println(a, b)
}
