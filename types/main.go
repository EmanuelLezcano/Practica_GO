package main

import "fmt"

type course struct {
	name string
}

func (c course) Print() {
	fmt.Printf("%+v\n", c)
}

//Declaración de alias
type myAlias = course

/* Cuando defino un nuevo tipo solo puedo acceder por ej a los campos del struct en este caso type course, pero no a los métodos, automaticamente se quitan todos los métodos.
NOTA: es totalmente diferente a cualquier otro tipo de dato, inclusive es diferente del tipo base sobre el que fué creado, en este caso course. */
//Definición de tipo
type newCourse course

/* Usualmente se ocupa para crear nuevos métodos para este nuevo tipo */
//Nuevo tipo en base a tipos ya establecidos
type newBool bool

func (b newBool) String() string {
	if b {
		return "VERDADERO"
	} else {
		return "FALSO"
	}
}
func main() {
	/* c := newCourse{name: "Go"} */
	/*
		---- Esto falla ----
		c.Print()
	*/
	/* fmt.Printf("El tipo es: %T\n", c) */
	var b newBool = true
	fmt.Println(b.String())
}
