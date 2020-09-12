package main

import "fmt"

type IObserver interface {
	Update()
}

type Observer struct {
	name string
}

func (o *Observer) Update() {
	fmt.Printf("%s触发\n", o.name)
}
