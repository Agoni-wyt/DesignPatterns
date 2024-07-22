package main

import (
	actualcombat "builderPattern/actualCombat"
	"fmt"
)

//客户端

func main() {
	MpvBuilder := actualcombat.GetBuilder("MPV")
	SuvBuilder := actualcombat.GetBuilder("SUV")

	director := actualcombat.NewDirector(MpvBuilder)
	mpvCar := director.BuildCar()
	fmt.Printf("MPV Car %#v",mpvCar)


	director.SetBuilder(SuvBuilder)
	suvCar := director.BuildCar()
	fmt.Printf("SUV Car %#v",suvCar)

}
