package main

import (
	af "./abstractFactory"
	"fmt"
)

func main() {
	//fruit := sf.GetFruit("apple")
	//fmt.Println(fruit.Fruit())
	//fruit2 := sf.GetFruit("banana")
	//fmt.Println(fruit2.Fruit())

	// abstract factory
	factory := af.GetFactory("wuhan")
	fruit := factory.ChooseFruit("apple")
	fmt.Println(fruit.Fruit())
	factory = af.GetFactory("hainan")
	fruit = factory.ChooseFruit("banana")
	fmt.Println(fruit.Fruit())
}
