package actualcombat

import "fmt"

type Mac struct {
}

func NewMac() *Mac {
	return &Mac{}
}

func (m *Mac) ConvertToUSB() {
	fmt.Println("Lightning 接入 mac 电脑")
}
