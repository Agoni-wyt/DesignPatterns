package main

import actualcombat "adapterPettern/actualCombat"

func main() {
	Client := &actualcombat.Client{}
	Mac := &actualcombat.Mac{}
	Windowns := &actualcombat.Windowns{}
	Client.InsertIntoComputer(Mac)
	//Client.InsertIntoComputer(Windowns) //err  actualcombat.Computer (missing method ConvertToUSB)
	//使用适配器
	Adapter := &actualcombat.Adapter{
		WindowsMachine: Windowns,
	}
	Client.InsertIntoComputer(Adapter)
}