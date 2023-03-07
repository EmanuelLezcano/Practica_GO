package main

import "fmt"

func main() {

	for i := 0; i < 10; i++ {
		fmt.Print(i)
	}
	fmt.Print("\n")
	// Otra implementación como el while
	i := 0
	for i < 10 {
		fmt.Print(i)
		i++
	}

	//Otra implementación, indefinidamente para cuando nesesito que procesos se escuchen constantemente
	/* j := 0
	for {
		fmt.Print(j)
		j++
	} */

	// For range devuelve una COPIA de cada elemento del array
	fmt.Print("\n")
	nums := []uint8{2, 4, 6, 8}
	/*
		for range nums {
			fmt.Println("Hola")
		} */
	// Retorna el indice y el valor de cada elemento
	/* for i, v := range nums {
		fmt.Println("Indice: ", i, "Valor: ", v)
	} */
	// Operando a cada elemento
	for i := range nums {
		nums[i] *= 2
	}
	fmt.Println(nums)

	//ITERANDO SOBREA MAP: de la misma manera que en un array
	sports := map[string]string{"basketball": "altura", "soccer": "habilidad"}

	for key, v := range sports {
		fmt.Println("Key: ", key, "Valor: ", v)
	}

	//ITERANDO SOBREA STRINGS: de la misma manera que en un array
	hello := "hello"
	for i, v := range hello {
		fmt.Println("índice: ", i, "Valor: ", string(v))
	}

}
