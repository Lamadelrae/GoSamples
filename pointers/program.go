package main

import (
	"fmt"

	"strconv"
)

type Person struct {
	Name string
	Age  int
}

func main() {
	var matthew Person
	matthew.Name = "Matthew"
	matthew.Age = 21

	fmt.Println("Age:" + strconv.Itoa(matthew.Age))
	changeAgeTo0(&matthew)
	fmt.Println("Age:" + strconv.Itoa(matthew.Age))
}

func changeAgeTo0(person *Person) {
	person.Age = 0
}
