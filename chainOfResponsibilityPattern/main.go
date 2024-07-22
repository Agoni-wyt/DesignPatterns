package main

import actualcombat "chainOfResponsibilityPattern/actualCombat"

func main() {
	// 收银台
	cashier := &actualcombat.Cashier{}
	// 设置下一个部门:药房-->收银台
	drugstore :=&actualcombat.Drugstore{}
	drugstore.SetNext(cashier)

	//设置下一个部门：病房-->药房
	clinic := &actualcombat.Clinic{}
	clinic.SetNext(drugstore)

	//设置下一个部门：门诊-->病房

	reception:=&actualcombat.Reception{}
	reception.SetNext(clinic)

	//门诊操作病人
	patient:=&actualcombat.Patient{Name:"张三"}
	reception.Operate(patient)
}