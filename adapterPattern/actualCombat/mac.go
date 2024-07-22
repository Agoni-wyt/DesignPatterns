package actualcombat

import "fmt"

type Mac struct {
}

func (m *Mac) ConvertToUSB() {
	fmt.Println("Lightning 接入 mac 电脑")
}