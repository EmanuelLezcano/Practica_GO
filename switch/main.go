package main

import "fmt"

func main() {
	emoji := "cat"
	switch emoji {
	case "dog":
		fmt.Printf("Es un animal.")
	case "cat":
		fmt.Printf("Es un animal.")
	case "apple":
		fmt.Printf("Es una fruta.")
	case "banan":
		fmt.Printf("Es una fruta.")
	default:
		fmt.Println("No tiene coincidencias.")
	}
	//Otro metodo
	switch emoji {
	case "dog", "cat":
		fmt.Printf("Es un animal.")
	case "apple", "banan":
		fmt.Printf("Es una fruta.")
	default:
		fmt.Println("No tiene coincidencias.")
	}
	//Otro metodo
	switch {
	case emoji == "dog" || emoji == "cat":
		fmt.Printf("Es un animal.")
	case emoji == "apple" || emoji == "banan":
		fmt.Printf("Es una fruta.")
	default:
		fmt.Println("No tiene coincidencias.")
	}

}
