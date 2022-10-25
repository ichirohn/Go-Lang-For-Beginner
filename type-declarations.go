package main

import "fmt"

func main() {
	type NoKTP string
	type Married bool

	var NoKTPIchiro NoKTP = "1234567890"
	var MarriedStatus Married = true
	fmt.Println(NoKTPIchiro)
	fmt.Println(MarriedStatus)

}
