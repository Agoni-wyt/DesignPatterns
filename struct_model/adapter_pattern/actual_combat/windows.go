package actualcombat

import "fmt"

type Windowns struct {
}

func NewWindowns() *Windowns {
	return &Windowns{}
}
func (w *Windowns) InsertIntoUSB() {
	fmt.Println("USB 接入 Windows 电脑")
}
