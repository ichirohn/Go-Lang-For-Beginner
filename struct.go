package main

import "fmt"

type Customer struct {
	Name, Address string
	Age           int
}

func main() {
	var ichiro Customer
	ichiro.Name = "Muhammad Ichiro"
	ichiro.Address = "Indonesia"
	ichiro.Age = 23

	fmt.Println(ichiro.Name)
	fmt.Println(ichiro.Age)
	fmt.Println(ichiro.Age)

	Budi := Customer{
		Name:    "Budi",
		Address: "Tangerang",
		Age:     35,
	}

	fmt.Println(Budi)

	Anwar := Customer{"Amwar", "Tangerang", 35}
	fmt.Println(Anwar)

}
