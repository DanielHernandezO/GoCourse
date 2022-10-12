package main

import "fmt"

type person struct {
	name     string
	lastName string
}

func main() {
	person := person{name: "Daniel", lastName: "Hernandez"}
	person.updateName("Andres")
	fmt.Printf("%+v\n", person)
}

func (p *person) updateName(newName string) {
	*&p.name = newName
}
