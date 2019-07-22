package main

import (
	"fmt"
)

type ICar interface { //普通汽车
	reduceGas()
}

type IECar interface { //电动汽车
	reducePower()
}

type Car struct {
	Name  string
	Price uint64
}

func (car *Car) reduceGas() {
	fmt.Println("this is Gas Car")
}

func (car *Car) reducePower() {
	fmt.Println("this is Electric Car")
}

func main() {
	var gasCar ICar = new(Car) //这是一辆普通汽油汽车
	var eCar IECar = new(Car)  //这是一辆电动汽车
	var car Car                //这是一辆混合动力汽车
	gasCar.reduceGas()
	eCar.reducePower()
	car.reduceGas()
	car.reducePower()
}
