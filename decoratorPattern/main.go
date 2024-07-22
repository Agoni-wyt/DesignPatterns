package main

import (
	actualcombat "decoratorPattern/actualCombat"
	"fmt"
)

func main() {
	baseParts := &actualcombat.BaseParts{}
	partA := &actualcombat.BasePartsA{}

	iphone := &actualcombat.IPhone{}
	iphone.SetComponent(baseParts)
	fmt.Println("iphone price is ",iphone.GetAllPrice())

	xiaomi := &actualcombat.Xiaomi{}

	xiaomi.SetComponent(partA)
	fmt.Println("xiaomi price is ",xiaomi.GetAllPrice())
}

