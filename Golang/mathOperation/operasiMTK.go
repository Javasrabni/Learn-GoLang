package main

import "fmt"

func main() {
	// Augmented Assigment
	a := 1
	b := 5
	c := 10

	b -= 5
	fmt.Println("Augmented Assigment: ", a + b + c)

	// Unary Operation
	a = 5
	b = 10

	a++ 
	// fmt.Println(a)

	a--
	fmt.Println("Unary Operation: ", a)

	// Perbandingan 
	type compare string
	type result bool

	var ucup compare = "OKE"
	var ganteng compare = "ucup"

	var results result = ucup != ganteng
	fmt.Println(results)
}