package actualcombat

import "errors"

//定义一个接口，包含制作手机和制作电脑两个功能

type InterFaceProductFactory interface {
	MakePhone() InterfacePhone
	MakeComputer() InterfaceComputer
}

type InterfacePhone interface {
	SetColor(color string)
	SetSize(size int)
	GetColor() string
	GetSize() int
}


type InterfaceComputer interface{
	SetColor(color string)
	SetSize(size int)
	GetColor() string
	GetSize() int
}



func GetProductFactory(brand string)(InterFaceProductFactory,error){
	if brand == "xiaomi"{
		return &XiaomiFactory{},nil
	}
	if brand =="Lenovo"{
		return &LenovoFactory{},nil
	}
	return nil,errors.New("brand is err")
}


