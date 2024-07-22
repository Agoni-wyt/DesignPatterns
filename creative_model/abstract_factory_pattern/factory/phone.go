package factory

type IPhone interface {
	SetColor(color string)
	SetSize(size int)
	GetColor() string
	GetSize() int
}

type Phone struct {
	color string
	size  int
}

func (p *Phone) SetColor(color string) {
	p.color = color
}

func (p *Phone) SetSize(size int) {
	p.size = size
}

func (p *Phone) GetColor() string {
	return p.color
}

func (p *Phone) GetSize() int {
	return p.size
}
