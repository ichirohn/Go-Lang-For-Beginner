package main

import "fmt"

type Customer struct {
	Name, Address string
	Age           int
}

func (customer Customer) SayHi(name string) {
	fmt.Println("Hello, Yayang, My name is", customer.Name)
}

func main() {
	var ichiro Customer
	ichiro.Name = "Ichiro"
	ichiro.Address = "Sweden"
	ichiro.Age = 23

	ichiro.SayHi("Yayang")
}
