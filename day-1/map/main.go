package main

import "fmt"

func main() {
	// MAP
	/*
		map adalah tipe data yang memiliki key dan value
			{"name": "Aku", "age": 21}
	*/
	map1 := map[string]string{}
	map1["name"] = "Saya"

	// tipe data key dan value harus sama dengan yang dideklarasi
	// map1["umur"] = 21 --> error
	map1["umur"] = "21" // tidak error karena string
	map1["name"] = "Ryan"

	// cara ambil value dari sebuah map key
	name := map1["name"]

	fmt.Println("name", name)

	// kita juga bisa mengecek apakah key ada atau tidak
	// return value pertama adalah value, return value kedua adalah bool (true jika ada key, false jika tidak ada)
	umur, exists := map1["umur"]
	fmt.Println("umur", umur, "exists", exists)

	// delete adalah function yang digunakan untuk menghapus key dari map
	delete(map1, "name")

	// iterasi map dengan for range
	for key, v := range map1 {
		fmt.Println("key", key, "value", v)
	}

	// make adalah function yang digunakan untuk membuat map
	map2 := make(map[string]int, 5)
	fmt.Println("map 1", map1, " map 2", map2)

	// make adalah function yang digunakan untuk membuat map di dalam array
	mapInArr := make([]map[string]any, 0)
	mapInArr = append(mapInArr, map[string]any{
		"name": "Agus",
		"age":  21,
	})
	mapInArr = append(mapInArr, map[string]any{
		"name": "Budi",
		"age":  22,
	})

	fmt.Println("mapInArr", mapInArr)
}
