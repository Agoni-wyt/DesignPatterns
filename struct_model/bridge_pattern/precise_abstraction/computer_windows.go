package precise_abstraction

import (
	"bridge_pattern/abstract"
	"fmt"
)

type Windows struct {
	printer abstract.Printer
}

func (w *Windows) Print() {
	fmt.Println("Print request for windows")
	w.printer.PrintFile()
}

func (w *Windows) SetPrinter(p abstract.Printer) {
	w.printer = p
}
