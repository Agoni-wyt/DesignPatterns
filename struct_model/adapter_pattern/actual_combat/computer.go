package actualcombat

import "fmt"

// Client 适配器
type Client struct {
}

func NewClient() *Client {
	return &Client{}
}

// InsertIntoComputer 插入电脑（电脑）
func (c *Client) InsertIntoComputer(com Computer) {
	fmt.Println("client 接入 计算机")
	com.ConvertToUSB()
}

type Computer interface {
	ConvertToUSB()
}
