package main

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
