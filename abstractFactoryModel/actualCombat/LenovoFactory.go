package actualcombat

type LenovoFactory struct {
}

type LenovoPhone struct {
	Phone
}
type LenovoComputer struct {
	Computer
}

func (l *LenovoFactory) MakePhone() InterfacePhone {
	return &LenovoPhone{
		Phone: Phone{
			color: "black",
			size:5,
		},
	}
}

func (l *LenovoFactory) MakeComputer() InterfaceComputer {
	return &LenovoComputer{
		Computer: Computer{
			color: "black",
			size:5,
		},
	}
}

