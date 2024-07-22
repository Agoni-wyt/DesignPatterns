package builder

//生成器接口

type IBuilder interface {
	SetSeatsType()
	SetEngineType()
	SetNumber()
	GetCar() Car
}

func GetBuilder(builderType string) IBuilder {
	if builderType == "SUV" {
		return &SUVBuilder{}
	} else if builderType == "MPV" {
		return &MPVBuilder{}
	}
	return nil
}
