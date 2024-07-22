package gun

import gunfactory "factory_method/gun_factory"

type Ak47 struct {
	gunfactory.Gun
}

func NewAk47() gunfactory.IGun {
	return &Ak47{
		gunfactory.Gun{
			Name:  "AK47 gun",
			Power: 4,
		},
	}
}
