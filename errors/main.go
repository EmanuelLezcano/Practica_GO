package main

import (
	"errors"
	"fmt"
	"io/ioutil"
)

func main() {
	//ERROR al leer un archivo
	contetn, err := ioutil.ReadFile("./hello.txt")
	if err != nil {
		fmt.Printf("Ocurrio un error: %v", err)
		return
	}
	fmt.Println(string(contetn))

	result, err := division(10, 0)
	if err != nil {
		fmt.Printf("Ocurrio un error: %v", err)
		return
	}
	fmt.Print(result)

}

/* func division(dividendo, divisor float32) (float32, error) {
	if divisor == 0 {
		return 0, errors.New("No puedes dividir por cero.")
	}
	return dividendo / divisor, nil
} */
//RETORNAR PARAMETROS NOMBRADOS (no recomendable)
func division(dividendo, divisor float32) (result float32, err error) {
	if divisor == 0 {
		err = errors.New("No puedes dividir por cero.")
		return
	}
	result = dividendo / divisor
	return
}
