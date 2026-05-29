package main 

import "fmt"

func main() {

	// Array Cara 1
	var array [5]string
	array[0] = "1"
	array[1] = "2"
	array[2] = "3"
	array[3] = "4"
	array[4] = "5"

	fmt.Println(array)

	// Array Cara 2
	anomali := 20
	var array_2 = [2]int {1, (anomali + 10) -27}
	fmt.Println(array_2)

	X := 2
	var buah = [6]string {"apel", "pisang", "nanas", "jambu", "nangka"}
	fmt.Println(buah[len(buah) - X])

	X += 29 //Augmented Assigment
	X ++ //Unary Operation

	fmt.Println(X)
}