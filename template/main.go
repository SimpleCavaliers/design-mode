package main

import "fmt"

type Game interface {
	Initialize()
	StartPlay()
	EndPlay()
}

type Cricket struct {}

func (*Cricket) Initialize() {
	fmt.Println("Cricket Game Initialized! Start playing.")
}

func (*Cricket) StartPlay() {
	fmt.Println("Cricket Game Started. Enjoy the game!")
}

func (*Cricket) EndPlay() {
	fmt.Println("Cricket Game Finished!")
}

type Football struct {}

func (*Football) Initialize() {
	fmt.Println("Football Game Initialized! Start playing.")
}

func (*Football) StartPlay() {
	fmt.Println("Football Game Started. Enjoy the game!")
}

func (*Football) EndPlay() {
	fmt.Println("Football Game Finished!")
}

func Do(game Game){
	game.Initialize()
	game.StartPlay()
	game.EndPlay()
}

func main() {
	Do(&Football{})
	Do(&Cricket{})
}
