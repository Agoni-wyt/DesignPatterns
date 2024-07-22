package actualcombat

import "time"

//部门抽象类 Department == Reception
type Department interface{
	Operate(*Patient)
	SetNext(Department)
}

//前台类
type Reception struct{
	next Department
}
func (r *Reception)Operate(p *Patient){
	if p.RegistrationDone{
		r.next.Operate(p)
		return
	}else{
		println("Patient 正在登记")
		time.Sleep(time.Second)
		p.RegistrationDone = true
		r.next.Operate(p)
	}
}
func(r *Reception)SetNext(next Department){
	r.next = next
}


//病房
type Clinic struct{
	next Department
}
func(d *Clinic)Operate(p *Patient){
	if p.ClinicCheckUpDone{
		d.next.Operate(p)
		return
	}else{
		println("Patient 正在看病")
		time.Sleep(time.Second)
		p.ClinicCheckUpDone = true
		d.next.Operate(p)
	}
}
func(d *Clinic)SetNext(next Department){
	d.next = next
}


//药房
type Drugstore struct{
	next Department
}
func (m *Drugstore)Operate(p *Patient){
	if p.DrugstoreDone{
		m.next.Operate(p)
		return
	}else{
		println("patient 正在拿药")
		time.Sleep(time.Second)
		p.DrugstoreDone = true
		m.next.Operate(p)
	}
}
func (m *Drugstore)SetNext(next Department){
	m.next = next
}

//收银台
type Cashier struct{
	next Department
}
func(c *Cashier)Operate(p *Patient){
	if p.PaymentDone{
		println("patient 支付完成")
		return
	}else{
		println("Patient 正在支付")
		time.Sleep(time.Second)
		p.PaymentDone = true
		println("patient 支付完成")

	}
}
func (c *Cashier)SetNext(next Department){
	c.next = next
}