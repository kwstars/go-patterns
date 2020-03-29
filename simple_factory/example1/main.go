package main

import "fmt"

type Factory interface {
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

func CreateFactory(fruit string) (factory Factory) {
	//根据不同的语言种类，实例化不同的翻译类
	switch fruit {
	case "apple":
		factory = new(Apple)
	case "banana":
		factory = new(Banana)
	default:
		panic("no such translator")
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
}
