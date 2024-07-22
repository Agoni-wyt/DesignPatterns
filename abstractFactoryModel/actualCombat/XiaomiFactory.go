package actualcombat


type XiaomiFactory struct {
}

type XiaomiPhone struct {
	Phone
}
type XiaomiComputer struct {
	Computer
}

func (l *XiaomiFactory) MakePhone() InterfacePhone {
	return &XiaomiPhone{
		Phone: Phone{
			color: "White",
			size:5,
		},
	}
}

func (l *XiaomiFactory) MakeComputer() InterfaceComputer {
	return &XiaomiComputer{
		Computer: Computer{
			color: "White",
			size:5,
		},
	}
}