package main

import "fmt"

func main() {
	// a := [5]int{1, 2, 3, 2, 2}
	// // b := [...]int{1, 2, 3, 2}
	// var c [4]string
	// fmt.Printf("%v %T\n", c[0], c[0])
	// c[0] = "123"
	// c[2] = "sad"
	// c[1] = "das"
	// fmt.Printf("%v %T\n", c[0], c[0])
	// c[0] = "Messi"
	// fmt.Printf("%v %T\n", c[0], c[0])
	// // fmt.Printf("%v %T\n", a, a)
	// // fmt.Printf("%v %T\n", a[1], a[1])
	// // fmt.Printf("%v %T\n", b, b)
	// // fmt.Printf("%v %T\n", b[1], b[1])
	// // fmt.Printf("%v %T\n", c, c)

	// fmt.Println(len(c))

	// a := [3][3]int{{1, 0, 1}, {0, 1, 0}, {0, 0, 1}}
	// var b [3][3]int
	// b[0] = [3]int{1, 2, 2}
	// b[1] = [3]int{2, 2, 2}
	// b[2] = [3]int{4, 2, 2}
	// fmt.Println(a)
	// fmt.Println(b)

	// a := [4]int{1, 3, 4, 1}
	// b := a
	// b[2] = 123
	// fmt.Println(a, b)

	// a := [...]int{123, 312, 13}
	// b := &a
	// b[1] = 2
	// fmt.Println(a, b)

	// a := []int{21312, 312, 31, 23, 123, 12, 3, 123}
	// // b := a
	// // b[1] = 222
	// // fmt.Println(a)
	// // fmt.Println(len(a))
	// // fmt.Println(len(b))
	// // fmt.Println(b)
	// fmt.Println(cap(a))
	// b := a[3:]
	// c := a[:6]
	// d := a[3:6]
	// e := a[:]
	// a[5] = 1
	// fmt.Println(b)
	// fmt.Println(c)
	// fmt.Println(d)
	// fmt.Println(e)
	// fmt.Println(cap(e))

	// a := make([]int, 3, 100)
	// fmt.Println(cap(a))
	// fmt.Println(len(a))
	// a = append(a, 2, 1, 1)
	// fmt.Println(cap(a))
	// fmt.Println(len(a))
	// fmt.Print(a)

	a := []int{}
	fmt.Println(a)
	fmt.Println(len(a))
	fmt.Println(cap(a))
	a = append(a, 1)
	fmt.Println(a)
	fmt.Println(len(a))
	fmt.Println(cap(a))
	a = append(a, 1)
	fmt.Println(a)
	fmt.Println(len(a))
	fmt.Println(cap(a))
	a = append(a, 1)
	fmt.Println(a)
	fmt.Println(len(a))
	fmt.Println(cap(a))
	a = append(a, 1)
	fmt.Println(a)
	fmt.Println(len(a))
	fmt.Println(cap(a))
	a = append(a, 1)
	fmt.Println(a)
	fmt.Println(len(a))
	fmt.Println(cap(a))
	a = append(a, 1)
	fmt.Println(a)
	fmt.Println(len(a))
	fmt.Println(cap(a))
	a = append(a, 1)
	fmt.Println(a)
	fmt.Println(len(a))
	fmt.Println(cap(a))
	a = append(a, 1)
	fmt.Println(a)
	fmt.Println(len(a))
	fmt.Println(cap(a))
	a = append(a, 1)
	fmt.Println(a)
	fmt.Println(len(a))
	fmt.Println(cap(a))

	// a := []int{1, 2, 3, 4, 5}
	// fmt.Println(a)
	// b := append(a[:2], a[3:]...)
	// fmt.Println(a)
	// fmt.Println(b)
}
