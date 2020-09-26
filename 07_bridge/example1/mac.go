package main

import "fmt"

type mac struct {
	printer IPrinter
}

func (m *mac) print() {
	fmt.Println("Print request for mac")
	m.printer.printFile()
}

func (m *mac) setPrinter(p IPrinter) {
	m.printer = p
}
