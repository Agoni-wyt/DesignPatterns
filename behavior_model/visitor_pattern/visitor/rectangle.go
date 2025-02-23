package visitor

type Rectangle struct {
	l int
	b int
}

func NewRectangle(l, b int) *Rectangle {
	return &Rectangle{l: l, b: b}

}

func (t *Rectangle) Accept(v Visitor) {
	v.visitForrectangle(t)
}

func (t *Rectangle) GetType() string {
	return "rectangle"
}
