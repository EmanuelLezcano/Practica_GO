package main

import (
	"curso_go/composition/customer"
	"curso_go/composition/invoice"
	"curso_go/composition/invoiceitem"
	"fmt"
)

func main() {
	i := invoice.New(
		"Colombia",
		"Bogotá",
		customer.New("Emanuel", "san patricio 272", "76"),
		invoiceitem.NewItems(
			invoiceitem.New(1, "Curso de GO", 12.3),
			invoiceitem.New(2, "Curso de HTML", 10.3),
			invoiceitem.New(3, "Curso de Angular", 8.3),
		),
	)
	i.SetTotal()
	fmt.Printf("%+v", i)
}
