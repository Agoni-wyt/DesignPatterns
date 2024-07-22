package actualcombat

//生成器接口

type InterfaceBuilder interface {
	SetSeatsType()
	SetEngineType()
	SetNumber()
	GetCar() Car
}

func GetBuilder(builderType string) InterfaceBuilder {
	if builderType == "SUV" {
		return &SUVBuilder{}
	} else if builderType == "MPV" {
		return &MPVBuilder{}
	}
	return nil
}