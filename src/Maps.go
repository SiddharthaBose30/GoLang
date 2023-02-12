package main

import (
	"fmt"
)

func main() {
	map1 := make(map[string]int)
	map1 = map[string]int{
		"Messi":  30,
		"Neymar": 10,
		"Mbappe": 7,
	}
	fmt.Println(map1)
	map1["Messi"] = 10
	fmt.Println(map1)
	map1["Pedri"] = 11
	fmt.Println(map1)
	delete(map1, "Mbappe")
	fmt.Println(map1)
	pop, ok := map1["Mbappe"]
	fmt.Println(pop, ok)
	pop1, ok1 := map1["Messi"]
	fmt.Println(pop1, ok1)
	_, ok2 := map1["Mbappe"]
	fmt.Println(ok2)
	fmt.Println(len(map1))
	map2 := map1
	map2["Messi"] = 1
	fmt.Println(map2)
	fmt.Println(map1)
}
