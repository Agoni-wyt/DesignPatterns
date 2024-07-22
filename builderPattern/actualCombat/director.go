package actualcombat

type Director struct{
	Builder InterfaceBuilder
}

func NewDirector(builder InterfaceBuilder) *Director{
	return &Director{Builder: builder}
}
func (d *Director) SetBuilder(builder InterfaceBuilder){
	d.Builder = builder
}
func (d *Director) BuildCar() Car{
	d.Builder.SetSeatsType()
	d.Builder.SetEngineType()
	d.Builder.SetNumber()
	return d.Builder.GetCar()
}