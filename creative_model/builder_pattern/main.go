package main

import (
	"builderPattern/builder"
	"fmt"
)

//客户端

func main() {
	// 获取建造器
	MpvBuilder := builder.GetBuilder("MPV")
	SuvBuilder := builder.GetBuilder("SUV")
	// 创建指挥者
	director := builder.NewDirector(MpvBuilder)
	// 建造MPV车
	mpvCar := director.BuildCar()
	fmt.Printf("MPV Car %#v", mpvCar)
	//更换建造器
	director.SetBuilder(SuvBuilder)
	// 建造SUV车
	suvCar := director.BuildCar()
	fmt.Printf("SUV Car %#v", suvCar)

}
