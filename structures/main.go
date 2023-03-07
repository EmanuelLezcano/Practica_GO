package main

import "fmt"

func main() {
	type course struct {
		Name      string
		Professor string
		Country   string
	}
	db := course{
		Name:      "Emanuel",
		Professor: "Emanuel-Lezcano",
		Country:   "Argentina",
	}
	// Otra manera
	git := course{
		"Juan",
		"Juan-Lezcano",
		"Brazil",
	}
	// Para cargar algunos campos
	css := course{
		Professor: "Emanuel-Lezcano",
	}
	fmt.Printf("%+v\n, %+v\n,%+v\n", db, git, css)

	//Accediendo a cada uno de los campos
	fmt.Printf("%+v\n, %+v\n", db.Name, git.Name)

	//Almaceno un puntero a la bd
	p := &db
	//Modificando el valor con el uso del puntero
	p.Name = "Beto"
	fmt.Printf("%+v\n", db)
	fmt.Printf("%+v\n", p)
}
