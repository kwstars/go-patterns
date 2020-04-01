package main

import "fmt"

type windows struct{}

func (w *windows) insertInCirclePort() {
	fmt.Println("Insert circle port into windows machine")
}

type windowsAdapter struct {
	windowMachine *windows
}

func (w *windowsAdapter) insertInSquarePort() {
	w.windowMachine.insertInCirclePort()
}
