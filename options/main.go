package main

import "fmt"

type Options struct {
	A string
	B string
	C int
}

// 定义选项模式的默认这
var (
	defaultOption = &Options{
		A: "A",
		B: "B",
		C: 100,
	}
)

// 定义一个OptionFunc 函数类型
type OptionFunc func(*Options)

// 利用闭包为每个字段编写一个设置值的With函数
func WithA(a string) OptionFunc {
	return func(o *Options) {
		o.A = a
	}
}

func WithB(b string) OptionFunc {
	return func(o *Options) {
		o.B = b
	}
}

func WithC(c int) OptionFunc {
	return func(o *Options) {
		o.C = c
	}
}

func newOption(a, b string, c int) *Options {
	return &Options{
		A: a,
		B: b,
		C: c,
	}
}

func newOption2(opts ...OptionFunc) (opt *Options) {
	opt = defaultOption
	for _, o := range opts {
		o(opt)
	}
	return
}

func main() {
	x := newOption("test01", "xxxxx", 10)
	fmt.Println(x)

	// 默认选项
	x = newOption2()
	fmt.Println(x)

	// 新的Option选项
	x = newOption2(WithA("test02"), WithC(8888))
	fmt.Println(x)
}
