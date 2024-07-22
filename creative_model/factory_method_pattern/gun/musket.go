package gun

import gunfactory "factory_method/gun_factory"

type Musket struct {
	gunfactory.Gun
}

func NewMusket() gunfactory.IGun {
	return &Musket{
		gunfactory.Gun{
			Name:  "Musket gun",
			Power: 1,
		},
	}
}
