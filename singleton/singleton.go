package main

import (
	"fmt"
	"sync"
)

type singleton struct {}
var instance *singleton
var mu sync.Mutex
var once sync.Once
var instance2 = &singleton{}

func GetInstance() *singleton{
	// double check
	if instance == nil{
		mu.Lock()
		defer mu.Unlock()
		if instance == nil{
			instance = &singleton{}
		}
	}
	return instance
}

func GetInstance2() *singleton{
	once.Do(func() {
		instance = &singleton{}
	})
	return instance
}

func GetInstance3() *singleton{
	return instance2
}

func main() {
	count := 10
	var wg sync.WaitGroup
	wg.Add(count)
	for i:=0; i<count; i++{
		go func() {
			defer wg.Done()
			instance := GetInstance3()
			fmt.Printf("%p\n", instance)
		}()
	}
	wg.Wait()
}
