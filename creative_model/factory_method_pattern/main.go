package main

import (
	"factory_method/gun"
	"fmt"
)

func main() {
	// Create a gun factory
	ak47 := gun.NewAk47()
	musket := gun.NewMusket()

	//get gun attribute
	fmt.Printf("gun name is: %s ,power is %d \n", ak47.GetName(), ak47.GetPower())
	fmt.Printf("gun name is: %s ,power is %d \n", musket.GetName(), musket.GetPower())
}
