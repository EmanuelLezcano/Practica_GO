package main

import (
	"fmt"
	"os"
)

func main() {
	//defer , aplaza la ejecución de una función hasta que la ejecución de la actual haya culminado
	//USOS: Cerrar archivos, cerrar conexiones, etc.

	/* defer fmt.Println(3)
	fmt.Println(1)
	fmt.Println(2) */

	//En este caso los defer se almacenan en una pila, y luego se van ejecutando en el orden de la pila (1,2,3)
	defer fmt.Println(3)
	defer fmt.Println(2)
	defer fmt.Println(1)

	//NOTA: el defer solo aplaza la ejecucion pero la asignacion a variables re realiza en orden de ejecucion normal ej...

	a := 5
	defer fmt.Println("Defer: ", a)
	a = 10
	fmt.Println(a)

	//USOS
	//package de go para crear un archivo
	file, err := os.Create("helloword.txt")
	if err != nil {
		fmt.Printf("Ocurrió un error al crear el archivo: %v", err)
		return
	}
	defer file.Close()
	_, err = file.Write([]byte("Hello emanuel"))
	if err != nil {
		//file.Close()
		fmt.Printf("Ocurrió un error al escribir el archivo: %v", err)
		return
	}
	//file.Close()
}
