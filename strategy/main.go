package main

import "fmt"

type Strategy interface {
	DoOperation(num1, num2 int) int
}

type OperationAdd struct {}
type OperationSub struct {}
type OperationMul struct {}

func (*OperationAdd) DoOperation(num1, num2 int) int{
	return num1 + num2
}

func (*OperationSub) DoOperation(num1, num2 int) int{
	return num1 - num2
}

func (*OperationMul) DoOperation(num1, num2 int) int{
	return num1 * num2
}

type Context struct {
	strategy Strategy
}

func (c *Context) executeStrategy(num1, num2 int) int{
	return c.strategy.DoOperation(num1, num2)
}

func main() {
	context := &Context{strategy: &OperationAdd{}}
	fmt.Println(context.executeStrategy(10, 5))

	context = &Context{strategy: &OperationSub{}}
	fmt.Println(context.executeStrategy(10, 5))

	context = &Context{strategy: &OperationMul{}}
	fmt.Println(context.executeStrategy(10, 5))
}
