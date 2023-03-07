package main

import "fmt"

func main() {
	fruit := "apple"
	//Indico que p va ser un apuntador a un string
	var p *string
	p = &fruit
	fmt.Printf("\nTipo: %T, Valor: %s, Direction: %v", fruit, fruit, &fruit)
	//Con %v se visualizan las direcciones de memoria
	fmt.Printf("\nTipo: %T, Valor: %v, Desreferenciación: %s", p, p, *p)
}
