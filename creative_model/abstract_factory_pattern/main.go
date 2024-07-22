package main

import (
	"abstract_factory/factory"
	"fmt"
)

func main() {
	XiaomiFactory, _ := factory.GetProductFactory("xiaomi")

	Lenovo, _ := factory.GetProductFactory("lenovo")

	xiaomiPhone := XiaomiFactory.MakePhone()
	xiaomiComputer := XiaomiFactory.MakeComputer()

	LenovoPhone := Lenovo.MakePhone()
	LenovoComputer := Lenovo.MakeComputer()

	fmt.Printf("xiaomi phone color:%s size:%d\n", xiaomiPhone.GetColor(), xiaomiPhone.GetSize())
	fmt.Printf("xiaomi computer color:%s size:%d\n", xiaomiComputer.GetColor(), xiaomiComputer.GetSize())

	fmt.Printf("Lenovo phone color:%s\n size:%d", LenovoPhone.GetColor(), LenovoPhone.GetSize())
	fmt.Printf("Lenovo computer color:%s size:%d\n", LenovoComputer.GetColor(), LenovoComputer.GetSize())
}
