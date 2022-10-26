package main

import "fmt"

// operasi perbandingan

func main() {

	var name1 = "Ichiro"
	var name2 = "Ichiro"

	var result bool = name1 == name2
	fmt.Println(result)

	var value1 = 100
	var value2 = 200

	fmt.Println(value1 > value2)  // kurang dari
	fmt.Println(value1 < value2)  // lebih dari
	fmt.Println(value1 == value2) // sama dengan
	fmt.Println(value1 != value2) // tidak sama dengan

}
