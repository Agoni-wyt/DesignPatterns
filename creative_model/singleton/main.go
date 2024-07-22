package main

import (
	"fmt"
	"sync"
)

// 单例实例
type singleton struct {
	Value int
}

type Singleton interface {
	getValue() int
}

func (s singleton) getValue() int {
	return s.Value
}

var (
	instance *singleton
	once     sync.Once
)

// GetInstance 构造方法，用于获取单例模式对象
func GetInstance(v int) Singleton {
	once.Do(func() {
		instance = &singleton{Value: v}
	})

	return instance
}

func main() {
	ins1 := GetInstance(1)
	ins2 := GetInstance(2)
	if ins1 != ins2 {
		_ = fmt.Errorf("instance is not equal %d != %d", ins1.getValue(), ins2.getValue())
	}
}
