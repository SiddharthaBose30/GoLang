package main

import "fmt"

type Doctor struct {
	Name  int
	Actor string
	des   []string
}

func main() {

	s := Doctor{
		Name:  10,
		Actor: "Messi",
		des:   []string{"YES"},
	}
	fmt.Println(s)

	a := struct{ name string }{name: "Messi"}
	b := a
	b.name = "GOAT"
	c := &a
	c.name = "PP"
	fmt.Println(b)
	fmt.Println(a)
	fmt.Println(c)
}
