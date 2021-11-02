package main

import "fmt"

func main() {
	fmt.Println("Hello world")

	number := 5
	number2 := 81

	hasil := number * number2
	var tampil string
	if hasil%2 == 0 {
		tampil = "hasilnya adalah genap"
	} else {
		tampil = "hasilnya adalah ganjil"
	}
	fmt.Println(tampil)
}
