package main

import "fmt"

type Observer interface {
	Update()
}

type Subject struct {
	observers map[Observer]bool
}

func (s *Subject) AddObserver(observer Observer) {
	if _, ok := s.observers[observer]; !ok{
		s.observers[observer] = true
	}
}

func (s *Subject) SetState(){
	s.notifyAllObserver()
}

func (s *Subject) notifyAllObserver() {
	for observer := range s.observers{
		observer.Update()
	}
}

type ObserverA struct {}

func (*ObserverA) Update() {
	fmt.Println("A : new message")
}
type ObserverB struct {}

func (*ObserverB) Update() {
	fmt.Println("B : new message")
}

func main() {
	subject := &Subject{observers: map[Observer]bool{}}

	subject.AddObserver(&ObserverA{})
	subject.AddObserver(&ObserverB{})
	subject.SetState()
	subject.SetState()
}
