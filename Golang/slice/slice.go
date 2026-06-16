package main

import "fmt"

func main() {
	// Perbedaannya ada pada pendeklarasian panjang array atau valuenya jika array, secara ekspilsit dituliskan "[...]" jika slice tidak "[]"

	item := []string{"Mobile", "Desktop", "PC", "2", "2", }

	// Array
	array := [...]string{"x","y","z"}
	appendAr := array[:]
	appendArray := append(appendAr, item...)

	fmt.Println("Append array", appendArray)

	// Tipe data Slice
	slice := []string {
		"1",
		"2",
		"3",
	}
	var sliceAppend = make([]string, 5, 10) 
	sliceAppend = append(sliceAppend, slice...)
	fmt.Println("Panjang:", len(sliceAppend))
	fmt.Println("Kapasitas:", cap(sliceAppend))

	fmt.Println(array)
	fmt.Println(slice)
	fmt.Println(sliceAppend)
}