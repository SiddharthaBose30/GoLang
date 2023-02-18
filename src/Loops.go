package main

import "fmt"

func main() {
	for i, j := 0, 1; i < 10 && j < 4; i, j = i+1, j+1 {
		// fmt.Println(i * j)
	}
	k := 0
	for k < 12 {
		// println(k)
		k += 2
	}
	map1 := make(map[string]int)
	map1 = map[string]int{
		"Messi":  30,
		"Neymar": 10,
		"Mbappe": 7,
	}
	a := [5]int{1, 2, 3, 2, 2}
	b := []string{"1", "2", "3"}

	for k, v := range map1 {
		println(k, v)
	}
	for _, v := range a {
		println(v)
	}
	for k := range b {
		println(k)
	}
	c := "dasd "
	for k, v := range c {
		fmt.Printf("%v %v ", k, string(v))
	}
}
