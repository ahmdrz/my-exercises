package main

import (
	"fmt"
)

type vehicleFactory struct {
	color      string
	bodyPrice  int64
	wheelPrice int64
}

type vehicle struct {
	color string
	price int64
}

func newVehicle() vehicleFactory {
	return vehicleFactory{}
}

func (v vehicleFactory) setColor(c string) vehicleFactory {
	v.color = c
	return v
}

func (v vehicleFactory) setBodyPrice(p int64) vehicleFactory {
	v.bodyPrice = p
	return v
}

func (v vehicleFactory) setWheelPrice(p int64) vehicleFactory {
	v.wheelPrice = p
	return v
}

func (v vehicleFactory) build() vehicle {
	return vehicle{
		color: v.color,
		price: v.bodyPrice + v.wheelPrice,
	}
}

func main() {
	mySpecialCar := newVehicle().setColor("red").setWheelPrice(120).setBodyPrice(200).build()
	fmt.Println(mySpecialCar.price)
}
