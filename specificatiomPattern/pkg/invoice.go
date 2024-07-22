package pkg

// 检测目标对象：发票数据对象和要检测的数据
type Invoice struct {
	Day    int
	Noice  int
	IsSend bool
}

// 抽象规格接口
type Specification interface {
	IsSatisfiedBy(invoice Invoice) bool
	And(Specification) Specification
	Or(Specification) Specification
	Not() Specification
	Relate(Specification)
}


// 发票数据到期规格类

type OverDueSpecification struct {
	Specification
}
func (os *OverDueSpecification)IsSatisfiedBy(in Invoice)bool{
	return in.Day >= 30
}
//发票到期规格
func NewOverDueSpecification()Specification{
	a:= &OverDueSpecification{&CompositeSpecification{}}
	a.Relate(a)
	return a
}


type NoiceSentSpecification struct{
	Specification
}
func(ns *NoiceSentSpecification)IsSatisfiedBy(in Invoice)bool{
	return in.Noice >= 3
}
//发票通知发送规格对象
func NewNoiceSentSpecification()Specification{
	a:=&NoiceSentSpecification{&CompositeSpecification{}}
	a.Relate(a)
	return a
}


//是否收到发票通知
type InCollectionSpecification struct{
	Specification
}
func(is *InCollectionSpecification)IsSatisfiedBy(in Invoice)bool{
	return !in.IsSend
}
func NewInCollectionSpecification()Specification{
	a:=&InCollectionSpecification{&CompositeSpecification{}}
	a.Relate(a)
	return a

}
