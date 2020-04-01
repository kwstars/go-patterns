package main

import "fmt"

type printer interface {
	printFile()
}

type epson struct{}

func (p *epson) printFile() {
	fmt.Println("Printing by a EPSON Printer")
}

type hp struct{}

func (p *hp) printFile() {
	fmt.Println("Printing by a HP Printer")
}

func main() {
	hpPrinter := &hp{}
	epsonPrinter := &epson{}
	macComputer := &mac{}
	macComputer.setPrinter(hpPrinter)
	macComputer.print()
	fmt.Println()
	macComputer.setPrinter(epsonPrinter)
	macComputer.print()
	fmt.Println()
	winComputer := &windows{}
	winComputer.setPrinter(hpPrinter)
	winComputer.print()
	fmt.Println()
	winComputer.setPrinter(epsonPrinter)
	winComputer.print()
	fmt.Println()
}
