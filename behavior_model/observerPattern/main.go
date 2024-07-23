package main

import "observerPattern/pkg"

func main() {
	bookProduct := pkg.NewProduct("book")
	observe1 := &pkg.User{Id: "qwe@gmail.com"}
	observe2 := &pkg.User{Id: "XXX@gmail.com"}

	bookProduct.Register(observe1)
	bookProduct.Register(observe2)

	bookProduct.UpdateAvailability()
}
