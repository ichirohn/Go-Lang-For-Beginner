package main

import "fmt"

func main() {

	var result = 10 * 10
	fmt.Println(result)

	var a = 1000
	var b = 5000
	var c = a + b
	fmt.Println(c)

	// augmented assigments

	var i = 10
	i += 10 // i = 1 + 10

	fmt.Println(i)

	// unary operator

	i++ // i = i + i
	fmt.Println(i)

	var negative = -100
	var possitive = 100
	fmt.Println(negative)
	fmt.Println(possitive)

}
