package factory

type LenovoFactory struct {
}

type LenovoPhone struct {
	Phone
}
type LenovoComputer struct {
	Computer
}

func (l *LenovoFactory) MakePhone() IPhone {
	return &LenovoPhone{
		Phone: Phone{
			color: "black",
			size:  5,
		},
	}
}

func (l *LenovoFactory) MakeComputer() IComputer {
	return &LenovoComputer{
		Computer: Computer{
			color: "black",
			size:  5,
		},
	}
}
