package actualcombat
// 组件，等待装饰器进行装配
type BaseParts struct {
}
func (b *BaseParts) GetPrice()float32{
	return 2000
}


type BasePartsA struct {
}
func (b *BasePartsA) GetPrice()float32{
	return 2200
}


type BasePartsB struct {
}
func (b *BasePartsB) GetPrice()float32{
	return 3400
}