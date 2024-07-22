package actualcombat

import "fmt"

type Adapter struct {
	WindowsMachine *Windowns
}

func (a *Adapter) ConvertToUSB() {
	fmt.Println("适配器将Lightning接口信号转换为USB信号")
	a.WindowsMachine.InsertIntoUSB()
}