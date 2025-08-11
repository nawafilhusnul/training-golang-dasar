package main

import "fmt"

func main() {
	var a int = 30
	var b int = 20

	// Operator Aritmatika : +, -, *, /, %
	// digunakan untuk operasi matematika
	sum := a + b
	subtraction := a - b
	division := float64(a) / float64(b)
	mod := a % b

	fmt.Println("30 = ", sum, "-10 = ", subtraction, "1.5 = ", division, "10 = ", mod)

	// Operator Perbandingan : >, <, >=, <=, ==, !=
	// digunakan untuk membandingkan nilai
	lebihBesar := a >= b
	lebihKecil := a <= b

	samaDengan := a == b
	tidakSamaDengan := a != b

	fmt.Println("lebih besar = ", lebihBesar, "lebih kecil = ", lebihKecil, "sama dengan = ", samaDengan, "tidak sama dengan = ", tidakSamaDengan)

	// Operator Logika : &&, ||, !
	// digunakan untuk logika
	operatorDan := samaDengan && lebihKecil
	operatorOr := samaDengan || lebihKecil
	operatorNot := !samaDengan

	fmt.Println("operator dan = ", operatorDan, "operator or = ", operatorOr, "operator not = ", operatorNot)
}
