package main

import "fmt"

func main() {
	a := 1
	b := 2
	c := 3

	if a > b {
		fmt.Println("a > b")

		// nested if
		if a > c {
			fmt.Println("a > c")
		}
	} else if b > a {
		fmt.Println("b > a ")
	} else {
		fmt.Println("semua salah")
	}

	switch a {
	case 10:
		fmt.Println("a = 10")
	case 20:
		fmt.Println("a = 20")
	default:
		fmt.Println("salah semua")
	}

	switch {
	case a < b:
		fmt.Println("a kurang dari b")
	case a > b:
		fmt.Println("a lebih besar dari b")
	default:
		fmt.Println("Salah semua")
	}
}
