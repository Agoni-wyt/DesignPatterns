package pkg

import "fmt"

// 主题接口
type Subject interface {
	Register(Observer Observer)   //注册
	Deregister(Observer Observer) //取消注册
	NotifyAll()                   //通知所有用户
}

// 具体主题：商品类
type Product struct {
	ObserverList []Observer
	name         string
	inStock      bool
}

func NewProduct(name string) *Product {
	return &Product{
		name: name,
	}
}

func (i *Product) UpdateAvailability() {
	fmt.Println("商品上架")
	i.inStock = true
	i.NotifyAll()
}

// 注册到对象列表
func (i *Product) Register(o Observer) {
	i.ObserverList = append(i.ObserverList, o)
}

// 从观察者对象列表删除
func (i *Product) Deregister(o Observer) {
	i.ObserverList = RemoveFromslice(i.ObserverList, o)
}

// 通知所有用户
func (i *Product) NotifyAll() {
	for _, observer := range i.ObserverList {
		observer.Update(i.name)
	}
}

func RemoveFromslice(ObserverList []Observer, ObserverToRemove Observer) []Observer {
	ObserverListLength := len(ObserverList)
	for i, observer := range ObserverList {
		if ObserverToRemove.GetID() == observer.GetID() {
			ObserverList[ObserverListLength-1], ObserverList[i] = ObserverList[i], ObserverList[ObserverListLength-1]
			return ObserverList[:ObserverListLength-1]
		}

	}
	return ObserverList
}
