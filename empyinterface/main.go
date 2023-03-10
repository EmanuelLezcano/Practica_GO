package main

import (
	"fmt"
	"strings"
)

type exapler interface {
	x()
}

/* Una interface vacía, la puede implementar cualquier objeto con cualquier method */

func wrrapper(i interface{}) {
	fmt.Printf("\nValor: %v", i, "\nTipo: %T", i)

	/* Type assertion te devuelve true y el valor del objeto que quieres saber de que tipo es, entre parentesis se coloca el tipo que queres validar */

	/* v, ok := i.(string)
	if ok {
		fmt.Printf("\t%s\n", strings.ToUpper(v))
	} */

	/* Type switches es mas cómodo que type acertions */
	switch v := i.(type) {
	case string:
		fmt.Printf("\t%s\n", strings.ToUpper(v))
	case int:
		fmt.Printf("\tNúmero %v\n", v*2)
	default:
		fmt.Printf("\nValor: %v", v, "\nTipo: %T", v)
	}
}

func main() {
	/* var e exapler

	e.x() */
	wrrapper(3)
	wrrapper(true)
	wrrapper("hola")
	wrrapper(2.5)
	wrrapper("Alvaro")
}

/* MAYOR FLEXIVILIDAD CUANDO NO CONOCEMOS EL TIPO QUE VAMOS A RECIBIR */
