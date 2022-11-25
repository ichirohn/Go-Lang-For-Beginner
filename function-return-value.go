package main

import "fmt"

func getHello(name string) string {
	if name == "" {
		return "Hello Guys"
	} else {
		return "Hello " + name
	}
}

func main() {
	result := getHello("Ichiro")
	fmt.Println(result)

	fmt.Println(getHello(""))
}
