package main

import "fmt"

type motorVehicle interface {
	Mileage() float32
}

type BMW struct {
	distance float32
	fuel     float32
}

type Audi struct {
	distance float32
	fuel     float32
}

func (b BMW) Mileage() float32 {
	return b.distance / b.fuel
}
func (a Audi) Mileage() float32 {
	return a.distance / a.fuel
}

func total_mileage(m []motorVehicle) float32 {
	tm := float32(0.0)
	for _, v := range m {
		tm = tm + v.Mileage()
	}
	return tm
}
func main() {
	b := BMW{
		distance: 123.1,
		fuel:     12.12,
	}
	a := Audi{
		distance: 123.1,
		fuel:     12.12,
	}
	m := []motorVehicle{b, a}
	fmt.Println(total_mileage(m))
	found("String")
	found(12)
	found(true)
	found(map[string]int{})
	found(BMW{})
	found([]int{1, 2, 3})
	found([...]int{1, 2, 3, 4, 5, 5})
}

type i interface{}

func found(i interface{}) {
	fmt.Printf("type %T, value %v\n", i, i)
}
