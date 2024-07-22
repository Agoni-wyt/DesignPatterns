package builder

type MPVBuilder struct {
	SeatsType  string
	EngineType string
	Number     int
}

func NewMPVBuilder() *MPVBuilder {
	return &MPVBuilder{}
}
func (m *MPVBuilder) SetSeatsType() {
	m.SeatsType = "MPV座椅"
}
func (m *MPVBuilder) SetEngineType() {
	m.EngineType = "MPV发动机"
}
func (m *MPVBuilder) SetNumber() {
	m.Number = 8
}

func (m *MPVBuilder) GetCar() Car {
	return Car{
		SeatsType:  m.SeatsType,
		EngineType: m.EngineType,
		Number:     m.Number,
	}
}
