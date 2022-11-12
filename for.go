package main

import "fmt"

func main() {
	counter := 1

	for counter <= 10 {
		fmt.Println("Perulangkan Ke", counter)
		counter++
	}

	for counter := 1; counter <= 10; counter++ { //cara cepatnya dan simple
		fmt.Println("Perulangkan Ke", counter)
	}

	slice := []string{"Muhammad", "Ichiro", "L"}

	for i := 0; i < len(slice); i++ {
		fmt.Println(slice[i])
	}
	for i, value := range slice {
		fmt.Println("Index", i, "=", value)
	}

	person := make(map[string]string)
	person["name"] = "Ichiro"
	person["tittle"] = "Hampir Jadi Programmer"

	for key, value := range person {
		fmt.Println(key, "=", value)

	}

}
