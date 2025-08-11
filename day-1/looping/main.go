package main

import "fmt"

func main() {
	// for loop classic
	for i := 0; i < 5; i++ {
		fmt.Println("i = ", i)
	}

	// for loop while
	i := 0
	for i < 5 {
		fmt.Println("i = ", i)
		i++
	}

	// infinite for loop
	// wajib ada break untuk menghentikan loop, kalau tidak akan infinite loop
	j := 0
	for {
		if j == 10 {
			break
		}

		j++

		if j%2 == 0 {
			continue
		}

		fmt.Println("j = ", j)

	}

	// for loop range
	// bisa digunakan untuk array, slice, map, string
	myName := "Agus"
	for idx, s := range myName {
		fmt.Println("index ", idx, " value ", string(s))
	}

	fmt.Println(string(myName[1]))
}
