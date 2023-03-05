package main

import "fmt"

type Hasname interface {
	GetName() string
}

func SayHello(hasName Hasname) {
	fmt.Println("Hello World", hasName.GetName())
}

type Person struct {
	Name string
}

func (person Person) GetName() string {
	return person.Name
}

type Animal struct {
	Name string
}

func (animal Animal) GetName() string {
	return animal.Name
}

func main() {
	var ichiro Person
	ichiro.Name = "Ichiro"

	SayHello(ichiro)

	var kucing Animal
	kucing.Name = "Kucing"

	SayHello(kucing)

}
