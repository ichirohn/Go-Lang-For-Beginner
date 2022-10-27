package main

import "fmt"

// membuat array secara manual
func main() {
	var names [4]string
	names[0] = "Ichiro"
	names[1] = "Muhammad"
	names[2] = "Joko"
	names[3] = "Wiranto"

	fmt.Println(names[0])
	fmt.Println(names[1])
	fmt.Println(names[2])
	fmt.Println(names[3])

	// membuat array secara langsung
	var values = [3]int{
		90,
		95,
		80,
	}
	fmt.Println(values)
	fmt.Println(values[0])
	fmt.Println(values[1])
	fmt.Println(values[2])

	// mengetahui panjang array
	fmt.Println(len(names))
	fmt.Println(len(values))

	var oke [15]string
	fmt.Println(len(oke))

}
