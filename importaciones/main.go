package main

import (
	"fmt"

	"github.com/google/uuid"
)

func main() {
	fmt.Println("Hi")
	Id := uuid.ClockSequence()
	fmt.Println(Id)
}
