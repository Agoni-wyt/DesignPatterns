package actualcombat

// 装饰器
type Decorator struct {
	Part Part
}

func (d *Decorator)SetComponent(p Part){
	d.Part = p
}
func(d *Decorator)GetAllPrice(){
	if d.Part!=nil{
		d.Part.GetPrice()
	}
}