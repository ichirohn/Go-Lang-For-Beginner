package main

import "fmt"

func main() {

	var Ujian = 80
	var Absensi = 75

	var lulusUjian = Ujian >= 80
	var lulusAbsensi = Absensi >= 80
	fmt.Println(lulusUjian)
	fmt.Println(lulusAbsensi)

	var lulus = lulusUjian && lulusAbsensi
	fmt.Println(lulus)

	// cara cepetnya
	fmt.Println(Ujian >= 80 && Absensi >= 80)

}
