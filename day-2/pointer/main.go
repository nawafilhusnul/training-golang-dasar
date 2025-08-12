package main

import "fmt"

// REVIEW POINTER

func main() {
	// deklarasi variable
	var age int
	fmt.Println(age)
	fmt.Println(&age)

	// deklarasi pointer
	var nama1 *string = new(string)

	// inisialisasi pointer
	*nama1 = "Sentus"
	fmt.Println("nama 1 ", *nama1)
}
