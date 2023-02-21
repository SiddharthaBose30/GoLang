package main

import "fmt"

type motorVehicle interface {
	Mileage() float32
}

type BMW struct {
	distance     float32
	fuel         float32
	averageSpeed float32
}

func (b BMW) AverageSpeed() float32 {
	return b.averageSpeed
}
func (b BMW) Mileage() float32 {
	return b.distance / b.fuel
}

func avg_speed(m motorVehicle) float32 {
	u := m.(BMW)
	return u.AverageSpeed()
}
func main() {
	b := BMW{
		distance:     123.1,
		fuel:         12.12,
		averageSpeed: 123.3,
	}
	fmt.Println(avg_speed(b))

}
