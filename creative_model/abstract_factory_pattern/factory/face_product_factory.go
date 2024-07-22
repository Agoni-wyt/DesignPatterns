package factory

import "errors"

type InterFaceProductFactory interface {
	MakePhone() IPhone
	MakeComputer() IComputer
}

func GetProductFactory(brand string) (InterFaceProductFactory, error) {
	if brand == "xiaomi" {
		return &XiaomiFactory{}, nil
	}
	if brand == "Lenovo" {
		return &LenovoFactory{}, nil
	}
	return nil, errors.New("brand is err")
}
