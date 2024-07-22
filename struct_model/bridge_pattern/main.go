package main

import (
	"bridge_pattern/precise_abstraction"
	"fmt"
)

func main() {
	// 两个打印机
	hpPrinter := &precise_abstraction.Hp{}
	epsonPrinter := &precise_abstraction.Epson{}
	// mac电脑
	macComputer := &precise_abstraction.Mac{}
	//mac电脑连接hp打印机
	macComputer.SetPrinter(hpPrinter)
	macComputer.Print() // 打印
	fmt.Println()
	// mac电脑连接epson打印机
	macComputer.SetPrinter(epsonPrinter)
	macComputer.Print() // 打印
	fmt.Println()
	// Windows电脑 。。。
	winComputer := &precise_abstraction.Windows{}

	winComputer.SetPrinter(hpPrinter)
	winComputer.Print()
	fmt.Println()

	winComputer.SetPrinter(epsonPrinter)
	winComputer.Print()
	fmt.Println()
}
