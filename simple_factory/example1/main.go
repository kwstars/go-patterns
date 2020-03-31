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

func CreateFactory(fruit string) (product Product) {
	//根据不同的语言种类，实例化不同的翻译类
	switch fruit {
	case "apple":
		product = new(Apple)
	case "banana":
		product = new(Banana)
	default:
		panic("Unknown product")
	}

	return
}

func main() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
		}
	}()

	apple := CreateFactory("apple")
	apple.Produce()

	banana := CreateFactory("banana")
	banana.Produce()

	unknow := CreateFactory("unknow")
	unknow.Produce()
}
