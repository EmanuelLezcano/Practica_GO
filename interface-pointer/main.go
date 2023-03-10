package main

import "fmt"

type Storager interface {
	Get() string
	Set(string)
}
type Person struct {
	Name string
}

/* Colocar receptores de tipo puntero en todos los metodos de una estructura es una buena práctica, eso es una vez que nesesito de esto en un metodo. Tambien podríamos devolver en el constructor un puntero para generalizar todo. */

func NewPerson(name string) *Person {
	return &Person{name}
}
func (p *Person) Get() string {
	return p.Name
}
func (p *Person) Set(name string) {
	p.Name = name
}
func Exec(s Storager, name string) {
	s.Set(name)
	fmt.Println(s.Get())
}

func main() {
	p := NewPerson("Emanuel")
	/* Paso un puntero porque los receptores de tipo punteros engloban tanto a la estructura como a los metodos de ésta */
	Exec(p, "Alvaro")
}
