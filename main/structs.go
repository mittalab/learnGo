package main

import (
	"fmt"
)

const myConst float64 = 65535
const myConst2 float64 = 1.6454

type car struct {
	var1 uint16 // min 0 max 65535
	var2 int16
}

func (c car) abc() float64 {
	return myConst*float64(c.var2)
}

func (c *car) abcP() {
	c.var1 = 200
}

func main() {
	var my_car car = car{1,2}
	fmt.Println(my_car)
	my_car.abcP()
	fmt.Println(my_car)
}