package main

import "fmt"

func main() {
	a := make(map[string]int)
	a = map[string]int{
		"Messi": 10,
	}
	if pop, ok := a["Messi"]; ok {
		fmt.Println(pop)
	}
	// fmt.Println(pop) --> pop only defined inside if statement
	x := 50
	y := 60
	fmt.Println(x <= y, x >= y, x == y, x != y, x < y, x > y)
	if x < 50 || y > 40 {
		println("yes")
	}
	if x < 50 && y > 40 {
		println("yes")
	}

	switch 2 {
	case 1, 2, 3:
		fmt.Println("YES")
	case 4, 6, 5:
		fmt.Println("NO")
	}

	switch i := 22 + 3; i {
	case 1, 2, 3:
		fmt.Println("YES")
	case 4, 6, 25:
		fmt.Println("NO")
	}

	j := 123
	switch {
	case j <= 12:
		fmt.Println("YES")
	case j >= 123:
		fmt.Println("NO")
	}

	//switch with type
	type u struct {
		Name string
		Age  int
	}
	var k interface{} = u{}
	switch k.(type) {
	case int:
		fmt.Println("int")
	case string:
		fmt.Println("string")
	case bool:
		fmt.Println("boolean")
	case [3]int:
		fmt.Println("[3]int")
	case u:
		fmt.Println("struct u")
	}

}
