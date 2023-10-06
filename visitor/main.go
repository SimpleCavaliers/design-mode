package main

import "fmt"

type ComputerPart interface {
	Accept(visitor ComputerPartVisitor)
}

type Monitor struct {}

func (monitor *Monitor) Accept(visitor ComputerPartVisitor) {
	visitor.Visit(monitor)
}

type Mouse struct {}

func (mouse *Mouse) Accept(visitor ComputerPartVisitor) {
	visitor.Visit(mouse)
}

type ComputerPartVisitor interface {
	Visit(cp ComputerPart)
}

type ComputerPartDisplayVisitor struct {}

func (cpdv *ComputerPartDisplayVisitor) Visit(cp ComputerPart){
	switch cp.(type) {
	case *Monitor:
		fmt.Println("Displaying Monitor.")
	case *Mouse:
		fmt.Println("Displaying Mouse.")
	}
}

func main() {
	computer := &ComputerPartDisplayVisitor{}
	computer.Visit(&Monitor{})
	computer.Visit(&Mouse{})
}
