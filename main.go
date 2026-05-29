package main

import "fmt"

func main () {
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




}