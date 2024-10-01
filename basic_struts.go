package main

import "fmt"

type Person struct {
     name string
     age int
}

func main() {
     person := Person{name: "John", age:25}
	 fmt.Printf("Name: %s, Age: %d\n", person.name, person.age)

}
