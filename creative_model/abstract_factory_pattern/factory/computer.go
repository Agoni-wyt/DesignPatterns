package factory

type IComputer interface {
	SetColor(color string)
	SetSize(size int)
	GetColor() string
	GetSize() int
}

type Computer struct {
	color string
	size  int
}

func (c *Computer) SetColor(color string) {
	c.color = color
}

func (c *Computer) SetSize(size int) {
	c.size = size
}

func (c *Computer) GetColor() string {
	return c.color
}

func (c *Computer) GetSize() int {
	return c.size
}
