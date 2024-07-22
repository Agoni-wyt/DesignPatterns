package main

import actualcombat "adapterPettern/actual_combat"

func main() {
	Client := actualcombat.NewClient()
	Mac := actualcombat.NewMac()
	Windowns := actualcombat.NewWindowns()

	//
	Client.InsertIntoComputer(Mac)
	// windows 没有实现 Computer 接口，所以不能直接插入
	//使用适配器
	Adapter := &actualcombat.Adapter{
		WindowsMachine: Windowns,
	}
	Client.InsertIntoComputer(Adapter)
}
