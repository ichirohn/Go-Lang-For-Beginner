package main

import "fmt"

//tipe data map

func main() {

	person := map[string]string{
		"name":    "ichiro",
		"address": "tangerang",
	}

	//menambah key map
	person["tittle"] = "Programmer"

	fmt.Println(person)
	fmt.Println(person["name"])
	fmt.Println(person["address"])
	fmt.Println(person["Programmer"])

	//make map function code
	var book map[string]string = make(map[string]string)
	book["tittle"] = "Belajar Inggris"
	book["author"] = "Ichiro"
	book["release"] = "2022"
	book["Ups"] = "Salah"

	fmt.Println(book)
	fmt.Println(len(book))

	delete(book, "Ups")

	fmt.Println(book)
	fmt.Println(len(book))
}
