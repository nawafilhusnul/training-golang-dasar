package main

import "fmt"

func main() {
	// ini adalah komentar
	/*
		ini adalah komentar multiple line
	*/

	// fmt adalah package yang digunakan untuk print dan formatting
	// fmt.Println adalah function yang digunakan untuk print dengan new line
	fmt.Println("Hello World!")

	// fmt.Printf adalah function yang digunakan untuk print tanpa new line
	// \n adalah new line
	fmt.Printf("Hello World! \n")

	// fmt.Sprintf adalah function yang digunakan untuk print ke variable
	// fmt.Sprintf("Hello World! \n")
	name := "Rico"
	res := fmt.Sprintf("Hello %s\n", name)
	fmt.Println(res)
}
