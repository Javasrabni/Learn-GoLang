package main

import "fmt"

func main() {
	nama := "Javas"
	getChar := nama[0]
	fmt.Println(nama)
	fmt.Println(getChar)
	fmt.Println(string(getChar))


	type NamaKTP string
	type NIK int

	var ucup NamaKTP = "UCUP"
	var NIKUcup NIK = 90

	fmt.Println(ucup, NIKUcup)


	// Augmented Assigment
	a := 1
	b := 2
	c:= 10

	b *= 2
	c -= 8
	fmt.Println(a + b - c)

	c += 8
	fmt.Println(a + b -c)

	// Unary operator
	a++ 
	fmt.Println(a)

	// Perbandingan 
	type compare string
	var nama1 compare= "ucup"
	var nama2 compare = nama1

	var result bool = nama1 == nama2
	fmt.Println(result)

}