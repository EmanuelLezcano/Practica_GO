package main

import "fmt"

func main() {
	//Map: para crear un map utilizo la funcion make, en donde indico que voy a crear un map, entre corchetes va el tipo de dato de la llave o clave, y luego el tipo de dato que almacena dicha claves
	animals := make(map[string]string)
	animals["cat"] = "gato"
	animals["dog"] = "perro"
	fmt.Println("Animals: ", animals)

	//Otra manera de hacerlo (Mapa literal)
	fruits := map[string]string{
		"apple": "manzana",
		"banan": "banana",
	}
	fmt.Println("Fruits: ", fruits)

	//Delete elements, delete(map, key)
	delete(fruits, "apple")
	fmt.Println("Fruits: ", fruits)

	// Recuperar valores del map. Si la llave no existe devuelve el valor 0 del tipo de dato que almacena el map para los valores.

	fmt.Println("Value: ", fruits["apple"])
	//Recuperar valores y con verificación si la clave existe o no
	//fruit, ok := animals["melon"]
	//fmt.Println("Value: , Exist: ", fruit, ok)
	_, ok := fruits["melon"]
	if !ok {
		fruits["melon"] = "melon"
	}
	fmt.Println(fruits)
}
