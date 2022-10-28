package main

import "fmt"

func main() {

	//percabangan go-lang
	var name = "Ichiro"

	if name == "Ichiro" {
		fmt.Println("Hello Ichiro")
	} else if name == "Lagampung" {
		fmt.Println("Hello Lagampung")
	} else if name == "Yayang" {
		fmt.Println("Hello Yayang")
	} else {
		fmt.Println("Hi, salam kenal")
	}

	//if short-statement
	if length := len(name); length > 5 {
		fmt.Println("Terlalu Panjang")
	} else {
		fmt.Println("Nama Sudah Benar")
	}

}
