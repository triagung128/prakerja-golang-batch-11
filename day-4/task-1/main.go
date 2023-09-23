package main

import "fmt"

func main() {
	car := Car{
		Type:   "Sedan",
		FuelIn: 10,
	}

	fmt.Printf("%.2f mill", car.calculateDistance())
}

type Car struct {
	Type   string
	FuelIn float64
}

func (car Car) calculateDistance() float64 {
	return car.FuelIn / 1.5
}
