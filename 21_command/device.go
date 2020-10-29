package main

import "fmt"

type Device interface {
	On()
	Off()
}

type tv struct {
	isRuning bool
}

func (t *tv) On() {
	t.isRuning = true
	fmt.Println("tv on")
}

func (t *tv) Off() {
	t.isRuning = false
	fmt.Println("tv off")
}
