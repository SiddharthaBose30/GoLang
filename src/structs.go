package main

import (
	"fmt"
	"reflect"
)

type Doctor struct {
	Name  int
	Actor string
	des   []string
}
type A struct {
	Name string
	Age  int
}
type B struct {
	A
	Place  string
	height int
}

type C struct {
	Name string `required:"100"`
	Age  int
}

func main() {

	// s := Doctor{
	// 	Name:  10,
	// 	Actor: "Messi",
	// 	des:   []string{"YES"},
	// }
	// fmt.Println(s)

	// a := struct{ name string }{name: "Messi"}
	// b := a
	// b.name = "GOAT"
	// c := &a
	// c.name = "PP"
	// fmt.Println(b)
	// fmt.Println(a)
	// fmt.Println(c)
	d := B{}
	d.Name = "Messi"
	d.Age = 13
	d.height = 12
	d.Place = "Barcelona"
	fmt.Println(d)
	e := B{
		A{Name: "Neymar", Age: 10},
		"Brazil",
		123,
	}
	fmt.Println(e)
	t := reflect.TypeOf(C{})
	field, _ := t.FieldByName("Name")
	fmt.Println(field.Tag)
	fmt.Println(field.Type)

}
