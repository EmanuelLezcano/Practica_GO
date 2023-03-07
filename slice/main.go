package main

import "fmt"

func main() {
	// Slice no poseen datos, son apuntadores a un Array.
	// Los Slices tienen tamaño y capacidad:
	//		len(array) retorna el tamaño del array.
	//		cap(array) retorna la capacidad del array.
	// Funcion append(array, elem): esto agrega un nuevo elemento al array dinamicamente
	// Lo que hace internamente append():
	//		Una vez superada la capacidad(cap()), crea un nuevo array duplicando el tamaño del anterior y en este añade los elementos ya existentes junto con los nuevos elementos.
	set := [4]string{"lion", "cat", "dog", "fish"}
	// con set[init:final] indico desde que índice hasta que índice quiero obtener los elementos del array, si no especifíco el principio o el final por defecto es el primero y el ultimo elemento respectivamente, si no especifíco ni el principio ni el final, se imprimiran todos los elementos.
	animals := set[0:4]
	fly := set[2:4]
	fly = append(fly, "hola")

	fly[0] = "change"
	fmt.Println("Animals: ", animals)
	fmt.Println("Fly: ", fly)
	fmt.Println("Fly: ", fly[0])
	fmt.Println("Fly tamaño: ", len(fly))
	fmt.Println("Animals: ", animals)
}
