package pkg

// 组合规格类
type CompositeSpecification struct {
	Specification
}

// 检查规格
func (cs *CompositeSpecification) IsSatisfiedBy(in Invoice)bool{
	return false
}
// 与规格有关
func (cs *CompositeSpecification) Relate(specification Specification){
	cs.Specification = specification
}


//定义规格类：与
type AndSpecification struct{
	Specification
	compare Specification
}
func (as *AndSpecification)IsSatisfiedBy(in Invoice)bool{
	return as.Specification.IsSatisfiedBy(in)&&as.compare.IsSatisfiedBy(in)
}


// 规格操作:与
func (cs *CompositeSpecification) And(specification Specification) Specification{
	a := &AndSpecification{
		cs.Specification,specification,
	}
	a.Relate(a)
	return a
}


//定义规格类：或
type OrSpecification struct{
	Specification
	compare Specification
}
func (os *OrSpecification)IsSatisfiedBy(in Invoice)bool{
	return os.Specification.IsSatisfiedBy(in) || os.compare.IsSatisfiedBy(in)
}


// 规格操作:或
func (cs *CompositeSpecification) Or(specification Specification) Specification{
	a := &OrSpecification{
		cs.Specification,specification,
	}
	a.Relate(a)
	return a
}


//定义规格类：非
type NotSpecification struct{
	Specification
}
func(ns *NotSpecification)IsSatisfiedBy(in Invoice)bool{
	return ns.Specification.IsSatisfiedBy(in)
}
// 规格操作:非
func (cs *CompositeSpecification) Not() Specification{
	a := &NotSpecification{
		cs.Specification,
	}
	a.Relate(a)
	return a
}

