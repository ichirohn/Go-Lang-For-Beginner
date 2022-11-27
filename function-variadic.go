package main

import "fmt"

func sumALL(numbers ...int) int {
	total := 0

	for _, value := range numbers {
		total += value
	}
	return total

}

func main() {
	total := sumALL(10, 90, 150, 50, 30)
	fmt.Println(total)

	slice := []int{15, 10, 20, 30, 100}
	total = sumALL(slice...)
	fmt.Println(total)
}
