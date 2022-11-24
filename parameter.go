package main

import "fmt"

func sayHiTo(firstname string, lastname string) {
	fmt.Println("hello", firstname, lastname)
}

func main() {
	firstname := "ichiro"
	sayHiTo(firstname, "muhammad")
	sayHiTo("Kevin", "Mitnick")
}
