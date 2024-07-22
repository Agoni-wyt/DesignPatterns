package gunfactory

type IGun interface {
	setName(name string)
	setPower(power int)
	GetName() string
	GetPower() int
}

type Gun struct {
	Name  string
	Power int
}

func (g *Gun) setName(name string) {
	g.Name = name
}

func (g *Gun) setPower(power int) {
	g.Power = power
}

func (g *Gun) GetName() string {
	return g.Name
}

func (g *Gun) GetPower() int {
	return g.Power
}
