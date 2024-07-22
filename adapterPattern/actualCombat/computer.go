package actualcombat

import "fmt"

type Client struct {
}

// 插入电脑（电脑）
func (c *Client) InsertIntoComputer(com Computer) {
	fmt.Println("client 接入 计算机")
	com.ConvertToUSB()
}

type Computer interface {
	ConvertToUSB()
}