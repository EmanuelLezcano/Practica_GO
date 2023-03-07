package main

import "fmt"

func main() {
	var emoji = "cactus"
	if emoji == "cactus" {
		fmt.Printf("Es: %s", emoji)
	} else if emoji == "carita" {
		fmt.Println("Es una carita.")
	} else {
		fmt.Println("No es carita ni cactus.")
	}
	fmt.Println("\nS")
	//Otra forma, solo se puede usar en lo que viva el if
	if carita := "carita"; carita == "cactus" {
		fmt.Println("Es un cactus.")
	} else if carita == "carita" {
		fmt.Printf("Es: %s", carita)
	} else {
		fmt.Println("No es carita ni cactus.")
	}
}
