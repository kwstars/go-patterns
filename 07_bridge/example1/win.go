package main

import "fmt"

type windows struct {
	printer IPrinter
}

func (w *windows) print() {
	fmt.Println("Print request for windows")
	w.printer.printFile()
}

func (w *windows) setPrinter(p IPrinter) {
	w.printer = p
}
