package main

import (
	"container/list"
	"fmt"
	"reflect"
	"strconv"
)

type Employee struct {
	Name         string
	Dept         string
	Salary       int
	Subordinates *list.List
}

func NewEmployee(name, dept string, salary int) *Employee {
	sub := list.New()
	return &Employee{
		Name:         name,
		Dept:         dept,
		Salary:       salary,
		Subordinates: sub,
	}
}

func (e *Employee) Add(emp *Employee) {
	e.Subordinates.PushBack(emp)
}

func (e *Employee) Remove(emp *Employee) {
	for i := e.Subordinates.Front(); i != nil; i = i.Next() {
		if reflect.DeepEqual(i.Value, emp) {
			e.Subordinates.Remove(i)
			fmt.Println("Yes")
		}
	}
}

func (e *Employee) GetSubordinates() *list.List {
	return e.Subordinates
}

func (e *Employee) ToString() string {
	return "[ Name: " + e.Name + ", dept: " + e.Dept + ", Salary: " + strconv.Itoa(e.Salary) + " ]"
}

func main() {
	ceo := NewEmployee("John","CEO", 30000)

	headSales := NewEmployee("Robert","Head Sales", 20000)
	headMarketing := NewEmployee("Michel","Head Marketing", 20000)

	clerk1 := NewEmployee("Laura","Marketing", 10000)
	clerk2 := NewEmployee("Bob","Marketing", 10000)

	salesExecutive1 := NewEmployee("Richard","Sales", 10000)
	salesExecutive2 := NewEmployee("Rob","Sales", 10000)

	ceo.Add(headSales)
	ceo.Add(headMarketing)
	headMarketing.Add(clerk1)
	headMarketing.Add(clerk2)
	headSales.Add(salesExecutive1)
	headSales.Add(salesExecutive2)

	fmt.Println(ceo.ToString())
	for i := ceo.Subordinates.Front(); i != nil; i = i.Next() {
		em := i.Value.(*Employee)
		fmt.Println(em.ToString())
		for j := i.Value.(*Employee).Subordinates.Front(); j != nil; j = j.Next() {
			em := j.Value.(*Employee)
			fmt.Println(em.ToString())
		}
	}
	ceo.Remove(headSales)
}
