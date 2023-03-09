package main

import (
	"fmt"

	"github.com/EmanuelLezcano/methods"
)

func main() {
	Gu := methods.NewCourse("nombre course", 15.54, true)
	Gu.CourseSetUserIDs([]uint{14, 54, 33})
	Gu.CourseSetClasses(map[uint]string{
		1: "Course basic GO",
		2: "Course intermediate GO",
		3: "Course avanced GO",
	})
	fmt.Printf("%v\n", Gu)
	Gu.CourseSetName("Course of Ema")
	fmt.Printf("%v\n", Gu)
	fmt.Println("\n", Gu.CourseName())
	Gi := methods.NewCourse("Nombre course", 0, true)
	fmt.Printf("%v\n", Gi)

}
