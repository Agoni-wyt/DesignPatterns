package actualcombat

import "fmt"

// Adapter 适配器，实现 原有接口，且携带新类
type Adapter struct {
	WindowsMachine *Windowns
}

func (a *Adapter) ConvertToUSB() {
	fmt.Println("适配器将Lightning接口信号转换为USB信号")
	a.WindowsMachine.InsertIntoUSB()
}
