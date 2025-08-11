package main

import "fmt"

func main() {
	// ARRAY
	// array adalah tipe data yang memiliki ukuran tetap, tidak akan bisa diubah ataupun ditambah
	// array bisa di deklarasi langsung dengan tipe data
	arr1 := [3]bool{}
	arr1[1] = true
	fmt.Println("arr ", arr1)

	// SLICE
	// slice adalah tipe data yang memiliki ukuran yang bisa berubah
	// slice bisa di deklarasi langsung dengan tipe data
	arr2 := []int{}

	// append adalah function yang digunakan untuk menambahkan data ke slice
	arr2 = append(arr2, 1)
	arr2 = append(arr2, 2)

	// append bisa menambahkan lebih dari 1 data
	arr2 = append(arr2, 2, 2, 3, 4, 5)

	// append bisa menambahkan data dari slice lain
	arr3 := []int{1, 2, 3, 4, 5}
	arr2 = append(arr2, arr3...)
	fmt.Println("arr2 ", arr2)

	// make adalah function yang digunakan untuk membuat slice
	arr4 := make([]int, 5, 6)
	arr4 = append(arr4, 1)

	// len adalah function yang digunakan untuk mendapatkan panjang slice
	// cap adalah function yang digunakan untuk mendapatkan kapasitas slice
	fmt.Println("value ", arr4, " len ", len(arr4), " cap ", cap(arr4))

	// append adalah function yang digunakan untuk menambahkan data ke slice dan akan melebihi kapasitas slice yang sudah dideklarasi
	// jika melebihi kapasitas slice, maka akan membuat slice baru dengan kapasitas yang lebih besar x2
	arr4 = append(arr4, 2)
	fmt.Println("value ", arr4, " len ", len(arr4), " cap ", cap(arr4))

	// copy adalah function yang digunakan untuk menyalin data dari slice ke slice lain
	arr5 := make([]int, 5)
	copy(arr5, arr4)
	fmt.Println("value ", arr5, " len ", len(arr5), " cap ", cap(arr5))

	// cara delete index tertentu
	arr6 := []int{1, 2, 3, 4, 5}
	deletedIdx := 2
	arr6 = append(arr6[:deletedIdx], arr6[deletedIdx+1:]...)
	fmt.Println("value ", arr6, " len ", len(arr6), " cap ", cap(arr6))
}
