package main

import "fmt"

type Product interface {
	Produce()
}

type ChinaApple struct{}

func (*ChinaApple) Produce() {
	fmt.Println("开始生产ChaiaApple")
}

type ChinaBanana struct{}

func (*ChinaBanana) Produce() {
	fmt.Println("开始生产ChinaBanana")
}

type JapanApple struct{}

func (*JapanApple) Produce() {
	fmt.Println("开始生产JapanApple")
}

type JapanBanana struct{}

func (*JapanBanana) Produce() {
	fmt.Println("开始生产JapanBanana")
}
