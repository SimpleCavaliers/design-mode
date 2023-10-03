package main

import (
	"fmt"
	"math/rand"
	"strconv"
)

type Shape interface {
	Draw()
}

type Circle struct {
	color 	string
	x 		int
	y 		int
	radius 	int
}

func (c *Circle) SetX(x int) {
	c.x = x
}

func (c *Circle) SetY(y int) {
	c.y = y
}

func (c *Circle) SetRadius(radius int) {
	c.radius = radius
}

func (c *Circle) Draw() {
	fmt.Println("Circle: Draw() [Color : " + c.color +
		", x : " + strconv.Itoa(c.x) + ", y : " + strconv.Itoa(c.y) +
		", radius : " + strconv.Itoa(c.radius))
}

// shape factory
var circleFactory = map[string]*Circle{}

func GetCircle(color string) Shape {
	if _, ok := circleFactory[color]; !ok{
		circleFactory[color] = &Circle{color: color}
		fmt.Println("Creating circle of color : " + color)
	}
	return circleFactory[color]
}

// test
func getRandomColor() string {
	colors := []string{"Red", "Green", "Blue", "White", "Black"}
	return colors[rand.Intn(len(colors))]
}

func getRandomX() int {
	return rand.Intn(100)
}

func getRandomY() int {
	return rand.Intn(100)
}

func main() {
	rand.Seed(42)
	for i:=0; i < 20; i++{
		circle := GetCircle(getRandomColor()).(*Circle)
		circle.SetX(getRandomX())
		circle.SetY(getRandomY())
		circle.SetRadius(100)
		circle.Draw()
	}
}
