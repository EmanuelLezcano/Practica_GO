package main

import "fmt"

//Interface Greeter
type Greeter interface {
	Greet()
}

//Interface Byer
type Byer interface {
	Bye()
}

//Embebido de interfaces: solo se puede realizar si las interfaces que voy a embeber tienen metodos con disjuntos
type GetterByer interface {
	Greeter
	Byer
}

//Struct Person
type Person struct {
	Name string
}

//Implentation Greeter interface for Person
func (p Person) Greet() {
	fmt.Printf("\nHi, my name is %s", p.Name)
}

//Implentation Byer interface for Person
func (p Person) Byer() {
	fmt.Printf("\nBye, my name is %s", p.Name)
}

//Type data Text
type Text string

//Implentation Greeter interface for Text
func (t Text) Greet() {
	fmt.Printf("\nHola soy %s", t)
}

//Implentation Byer interface for Text
func (t Text) Byer() {
	fmt.Printf("\nChau soy %s", t)
}

func GreetAll(gs ...Greeter) {
	for _, g := range gs {
		g.Greet()
		fmt.Printf("\t Valor: %v, Tipo: %T\n ", g, g)
	}
}
func ByerAll(bs ...Byer) {
	for _, b := range bs {
		b.Bye()
		fmt.Printf("\t Valor: %v, Tipo: %T\n ", b, b)
	}
}

func main() {
	/* var g Greeter = Person{"Emanuel"}
	var t Greeter = Text("Emanuel")
	g.Greet()
	t.Greet() */
	p := Person{Name: "Emanuel"}
	var t Text = "Juana"
	GreetAll(p, t)
}
