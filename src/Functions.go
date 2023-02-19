package main

import "fmt"

func main() {

	// var n1 *string
	// n2 := "1"
	// n1 = &n2
	// m1("Messi", "GOAT")
	// pass_by_value(n1)
	// pass_by_reference(n1)
	// fmt.Println(*n1)
	// sum("Messi", 1, 2, 3, 4, 5, 6)
	// s := sum1(1, 2, 3, 4, 5, 6)
	// s := sum2(1, 2, 3, 4, 5, 6)
	// s := sum3(1, 2, 3, 4, 5, 6)
	// println(s)
	// d, err := divide(1.0, 0.0)
	// if err != nil {
	// fmt.Println(err)
	// } else {
	// fmt.Println(d)
	// }
	// anonymous()
	A := greet{
		Greeting: "HI",
		A:        10,
	}
	A.f()
}
func m1(n1 string, n2 string) {
	fmt.Println(n1 + " is " + n2)
}
func pass_by_value(n1 string) {
	fmt.Println(n1)
	n1 = "2"
	fmt.Println(n1)
}
func pass_by_reference(n1 *string) {
	fmt.Println(*n1)
	*n1 = "2"
	fmt.Println(*n1)
}
func sum(n1 string, ms ...int) {
	for _, v := range ms {
		fmt.Println(v)
	}
	fmt.Println(n1)
}
func sum1(ms ...int) int {
	value := 0
	for _, v := range ms {
		value += v
	}
	return value
}
func sum2(ms ...int) *int {
	value := 0
	for _, v := range ms {
		value += v
	}
	return &value
}
func sum3(ms ...int) (value int) {
	for _, v := range ms {
		value += v
	}
	return
}
func divide(a, b float32) (float32, error) {
	if b == 0.0 {
		return 0.0, fmt.Errorf("Cannot divide by zero")
	}
	return a / b, nil
}

func anonymous() {
	func() {
		fmt.Println("Hi")
	}()

	//-------------------------

	f := func() {
		fmt.Println("YES")
	}
	f()

	//-------------------------

	var fun func() = func() {
		fmt.Println("NO")
	}
	fun()

	//-------------------------

	var divide func(float32, float32) (float32, error)
	divide = func(f1, f2 float32) (float32, error) {
		if f2 == 0.0 {
			return 0.0, fmt.Errorf("Cannot divide by zero")
		}
		return f1 / f2, nil
	}
	d, err := divide(1.0, 2.0)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(d)
	}

	//-------------------------

}

type greet struct {
	Greeting string
	A        int
}

func (a greet) f() {
	fmt.Println(a.A, a.Greeting)
}
