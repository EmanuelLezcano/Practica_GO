package main

import "fmt"

func main() {
	//panic(): nos permite entrar en panico, osea cortar la ejecución del programa
	division(10, 2)
	division(4, 0)
	division(4, 2)
}
func division(dividendo, divisor int) {
	defer func() {
		//recover devuelve el contenido del panic
		if r := recover(); r != nil {
			fmt.Println("Recuperandome... ", r)
		}
	}()
	validarDivisor(divisor)
	fmt.Println(dividendo / divisor)
}
func validarDivisor(divisor int) {
	if divisor == 0 {
		panic("panic!")
	}
}

//Cuando quiero recuperarme del panic() usamos una funcion recover()
