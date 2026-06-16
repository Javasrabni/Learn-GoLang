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

	var arrayTEs [3] int
	arrayTEs[2] = 3
	arrayTEs[1] = 1
	arrayTEs[0] = 2

	fmt.Println(arrayTEs)

	// Array kalo ga tau perkiraan jumlah datanya berapa, gunakan "[...]"
	var tesArray = [...] string {
		"mobil",
		"motor",
	}

	fmt.Println(tesArray)
	fmt.Println(len(tesArray))

	// SLICE
	dataArray := [...]string {"alek", "ucup", "santoso", "robert"}
	// fmt.Println(dataArray)

	slice1 := dataArray[:] //Get all Array value
	slice2 := dataArray[3:] //Get data dari batas low ke high
	slice3 := dataArray[:3] //Get dara dari batas high ke low
	fmt.Println(slice1, slice2, slice3)

	// nama := []string{dataArray[0],  dataArray[2]}
	// fmt.Println(nama)


	belajarlagi := []string {
		"1",
		"2",
		"3",
	}
	belajarlagi[0] = "satu"
	belajarlagi2 := append(belajarlagi, "EMPAT")
	fmt.Println(belajarlagi)
	fmt.Println(belajarlagi2)
}