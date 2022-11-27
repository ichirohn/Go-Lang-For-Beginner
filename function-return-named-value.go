package main

import "fmt"

func getfullName2() (firstName string, middleName string, lastName string) {
	firstName = "Muhammad"
	middleName = "Ichiro"
	lastName = "Lagampung"
	return
}

func main() {
	a, b, c := getfullName2()
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
}
