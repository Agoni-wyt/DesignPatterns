package actualcombat

//病人

type Patient struct {
	Name              string
	RegistrationDone  bool //门诊登记
	ClinicCheckUpDone bool //病房检查
	DrugstoreDone     bool //药房取药
	PaymentDone       bool //收银台支付
}
