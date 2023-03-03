package main

import "fmt"

func main() {
	//bool, string , numeric
	var a bool = true
	//Los string se declaran solo con comillas dobles, no simples
	var b string = "soy un string"
	//Los numeric pueden ser "uint" (+), "int" (+ o -), float (23.34)
	var c int = 34
	//Las comillas simples devuelven un resultado en base a la tabla UNICODE de caracteres

	// Para sumar datos de distintos tipo debo hacer la converción de datos, ejemplo uint16(varible).
	//NOTA: las variables no cambian su tipo en GO

	fmt.Printf("Tipo: %T , Valor: %v", a, a)
	fmt.Printf("Tipo: %T , Valor: %v", b, b)
	fmt.Printf("Tipo: %T , Valor: %v", c, c)
}
