package actualcombat


type IPhone struct{
	Decorator
}

func (i *IPhone)GetAllPrice() float32{
	return i.Part.GetPrice() + 6000
}


type Xiaomi struct {
	Decorator
}

func(x *Xiaomi)GetAllPrice() float32{
	return x.Part.GetPrice() + 1000
}