package main

import "fmt"

/*
 Factory method: es un metodo que nos permite crear instancias de objetos de un tipo determinado, encapsulando así la creación de dichos objetos
*/

type PayMethod interface {
	Pay()
}
type Paypal struct{}

func (p Paypal) Pay() {
	fmt.Println("Pagado con Paypal.")
}

type Cash struct{}

func (p Cash) Pay() {
	fmt.Println("Pagado con Efectivo.")
}

type CreditCard struct{}

func (p CreditCard) Pay() {
	fmt.Println("Pagado con CreditCard.")
}

func Factory(method uint) PayMethod {
	switch method {
	case 1:
		return Paypal{}
	case 2:
		return Cash{}
	case 3:
		return CreditCard{}
	default:
		return nil
	}
}

func main() {
	var method uint
	fmt.Println("Método de pago: ")
	fmt.Println("\n1: Paypal\n2: Cash\n3: CreditCard")
	_, err := fmt.Scanln(&method)
	if err != nil {
		panic("Debe digitar un método valido.")
	}
	if method > 3 {
		panic("Debe digitar un método valido.")
	}

	payMethod := Factory(method)
	payMethod.Pay()
}
