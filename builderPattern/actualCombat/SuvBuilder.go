package actualcombat

type SUVBuilder struct{
	SeatsType string
	EngineType string
	Number int
}

func NewSUVBuilder() *SUVBuilder{
	return &SUVBuilder{}
}
func (m *SUVBuilder) SetSeatsType(){
	m.SeatsType = "SUV座椅"
}
func (m *SUVBuilder) SetEngineType(){
	m.EngineType = "SUV发动机"
}
func (m *SUVBuilder)SetNumber(){
	m.Number = 6
}

func (m *SUVBuilder)GetCar() Car{
	return Car{
		SeatsType: m.SeatsType,
		EngineType: m.EngineType,
		Number: m.Number,
	}
}