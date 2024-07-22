package precise_abstraction

import (
	"bridge_pattern/abstract"
	"fmt"
)

type Mac struct {
	printer abstract.Printer
}

func (m *Mac) Print() {
	fmt.Println("Print request for mac")
	m.printer.PrintFile()
}

func (m *Mac) SetPrinter(p abstract.Printer) {
	m.printer = p
}
