package main

import "fmt"

type ProxyInterface interface {
	Display()
}

// 实体
type RealImage struct {
	fileName string
}

func (ri *RealImage) Display() {
	out := "Displaying " + ri.fileName
	fmt.Println(out)
}

func (ri *RealImage) LoadFromDisk(fileName string){
	out := "Loading " + fileName
	fmt.Println(out)
}

// 代理
type ProxyImage struct {
	fileName string
	realImage *RealImage
}

func (pi *ProxyImage) Display() {
	if pi.realImage == nil{
		pi.realImage = &RealImage{fileName: pi.fileName}
		pi.realImage.LoadFromDisk(pi.fileName)
	}
	pi.realImage.Display()
}

func main() {
	pi := &ProxyImage{fileName: "test1.jpg"}
	pi.Display()
	pi.Display()
	pi2 := &ProxyImage{fileName: "test2.jpg"}
	pi2.Display()
	pi2.Display()
}
