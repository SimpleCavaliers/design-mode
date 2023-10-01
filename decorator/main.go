package main

import "fmt"

// 抽象组件
type Shape interface {
	Draw()
}

// 抽象装饰器
type Decorator interface {
	Decorator(shape Shape)
}

// 具体组件
type Circle struct {}
type Rectangle struct {}

func (*Circle) Draw()  {
	fmt.Println("Shape: Circle")
}

func (*Rectangle) Draw(){
	fmt.Println("Shape: Rectangle")
}

// 具体装饰器
type RedDecorator struct {}

func (*RedDecorator) Decorator(shape Shape){
	fmt.Println("Border color: Red")
	shape.Draw()
}


func main(){
	circle := &Circle{}
	recrangle := &Rectangle{}
	redDecorator := &RedDecorator{}
	redDecorator.Decorator(circle)
	redDecorator.Decorator(recrangle)
}
