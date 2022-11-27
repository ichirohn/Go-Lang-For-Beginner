package main

import "fmt"

func getGoodBye(name string) string {
	return "Good Bye" + name
}

func main() {
	saygoodbye := getGoodBye

	result := saygoodbye("Ichiro")
	fmt.Println(result)
	fmt.Println(getGoodBye("Ichiro"))
}
