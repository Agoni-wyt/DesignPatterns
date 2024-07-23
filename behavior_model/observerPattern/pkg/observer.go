package pkg

import "fmt"

type Observer interface {
	Update(string)
	GetID() string
}

type User struct {
	Id string
}

// 更新
func (c *User) Update(ProductName string) {
	fmt.Println("新通知：！！商品上架")
}

func (c *User) GetID() string {
	return c.Id
}
