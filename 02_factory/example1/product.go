package main

import "fmt"

type Product interface {
	Produce()
}

type Apple struct{}

func (*Apple) Produce() {
	fmt.Println("开始生产苹果")
}

type Banana struct{}

func (*Banana) Produce() {
	fmt.Println("开始生产香蕉")
}
