package main

import (
	"fmt"
	"sync"
)

type singleTon struct {}
var instance *singleTon
var mu sync.Mutex
var once sync.Once
var instance2 = &singleTon{}

func GetInstance() *singleTon{
	// double check
	if instance == nil{
		mu.Lock()
		defer mu.Unlock()
		if instance == nil{
			instance = &singleTon{}
		}
	}
	return instance
}

func GetInstance2() *singleTon{
	once.Do(func() {
		instance = &singleTon{}
	})
	return instance
}

func GetInstance3() *singleTon{
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
