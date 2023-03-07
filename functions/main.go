package main

import "fmt"

func main() {
	hello()
	helloName("Emanuel", "Lezcano")

	//Parametros por valor
	toChange := "carita"
	changeValue(toChange)
	fmt.Println(toChange)
	//Parametros por referencia
	toChangeRef := "carita"
	change(&toChangeRef)
	fmt.Println(toChangeRef)

	//Return value
	tex := "Emanuel"
	fmt.Println(suma(2, 5))
	fmt.Println(convert(tex))

	//Funciones con argumentos
	nros := []int{1, 43, 54, 3, 12, 60}
	results := filter(nros, func(i int) bool {
		if i >= 13 {
			return true
		}
		return false
	})
	fmt.Println(results)

	// Funciones que retornan otra funcion
	nombre := helloUser("Emanuel") //Ejecuto o en este caso asigno la primera funcion
	fmt.Println(nombre())          //Ejecuto la segunda funcion

	//FUNCION VARIATICA
	//
	fmt.Println(sum(1, 3, 7, 5, 2))
}
func hello() {
	fmt.Println("Hello word!!")
}

func helloName(firstName string, lastName string) {
	fmt.Println("Hello ", firstName, lastName)
}

//POR parametros
func changeValue(value string) {
	value = "valor cambiado"
}
func change(value *string) {
	*value = "valor cambiado"
}

//Retornando valores
func suma(value1, value2 int) int {
	return value1 + value2
}

//retorna varios valores
func convert(text string) (string, string) {
	return "Ema", "nuel"
}

//Funciones con argumentos
func filter(nros []int, callback func(int) bool) []int {
	resultado := []int{}
	for _, v := range nros {
		if callback(v) {
			resultado = append(resultado, v)
		}
	}
	return resultado
}

//Funciones que retornan otra funcion
func helloUser(name string) func() string {

	return func() string {
		return "Hello " + name
	}
}

//Funciones variaticas, esta funcion recibe muchos enteros y los opero como un array
func sum(numbers ...int) int {
	total := 0
	for i := 0; i < len(numbers); i++ {
		total += numbers[i]
	}
	return total
}
