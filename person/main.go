package main

import "fmt"

type Person struct {
	Name string
	Age  uint
}

// Constructor Person
func NewPerson(name string, age uint) Person {
	return Person{name, age}
}

//Greet of Person
func (p Person) Greet() {
	fmt.Println("Hello")
}

/* Embeber Person en Employee, los campos y métodos del tipo persona ahora se convierten en campos y métodos del tipo Employee
NOTA: 	cuando se invoquen métodos de éste, al tipo que harán referecias estos métodos van a ser al tipo interno osea Person
*/
type Employee struct {
	Person
	Human
	Salary float64
}

//Constructor for Employee
func NewEmployee(name string, age uint, salary float64) Employee {
	return Employee{NewPerson(name, age), NewHuman(2, 4), salary}
}

//Net salary of Employee
func (e Employee) Payroll() float64 {
	return e.Salary * 0.90
}

//Greet of Employee
func (p Employee) Greet() {
	fmt.Println("Hello desde Employee")
}

//Struct Human
type Human struct {
	Age      uint
	Children uint
}

//Constructor for Human
func NewHuman(age, children uint) Human {
	return Human{age, children}
}

//Greet of Human
func (p Human) Greet() {
	fmt.Println("Hello desde Human")
}

func main() {
	e := NewEmployee("Emanuel", 30, 1000)
	/* Accediendo a los campos de Person con una instancia de Employee */
	fmt.Println("Name: ", e.Name, "\nAge: ", e.Person.Age, "\nSalary: ", e.Salary, "\nNet salary: ", e.Payroll())

	/* Ejemplo de la NO sobreescritura de metodos  */
	e.Person.Greet()
	e.Greet()
	/* e.Payroll() */

	/* PARA hacer referencia a un campo o método que está en varias estructuras usadas por Employe (Struct embedded), debo especificar de que struct quiero que me traiga el campo o el método. Por ej: e.Person.Age y e.Human.Age */

	println("Age of Human: ", e.Human.Age)
}
