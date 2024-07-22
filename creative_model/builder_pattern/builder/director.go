package builder

// Director 管理类
type Director struct {
	Builder IBuilder
}

func NewDirector(builder IBuilder) *Director {
	return &Director{Builder: builder}
}

func (d *Director) SetBuilder(builder IBuilder) {
	d.Builder = builder
}

// BuildCar 建造车
func (d *Director) BuildCar() Car {
	d.Builder.SetSeatsType()
	d.Builder.SetEngineType()
	d.Builder.SetNumber()
	return d.Builder.GetCar()
}
