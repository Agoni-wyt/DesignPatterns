package main

import (
	"fmt"
	"specificationPattern/pkg"
)

func main() {
	//声明发票数据到期规格
	overDue := pkg.NewOverDueSpecification()

	// 声明发票通知发送规格
	noticeSent := pkg.NewNoiceSentSpecification()
	// 声明是否收到发票通知规格
	inCollection := pkg.NewInCollectionSpecification()

	sendToCollection:=overDue.And(noticeSent).And(inCollection.Not()) //发送已收到通知需要保证发票未逾期，并且已经发送发票通知，以及收款机构未收到发票通知

	object := pkg.Invoice{
		Day: 32, //>=30  true
		Noice: 6,//>=3   true
		IsSend: false,//false
	}
	result:=sendToCollection.IsSatisfiedBy(object)
	fmt.Println(result)
}