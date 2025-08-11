package main

import "fmt"

// deklarasi untuk global private variable
// hanya bisa dipakai di package dalam folder ini
var umurPrivate int

// deklarasi untuk global public variable
// bisa dipakai di package dalam folder ini dan di package lain
var UmurPublic int

func main() {

	/*
		CARA DEKLARASI VARIABLE
	*/

	// 1. var namaVariabel tipeData = value
	var familyName string = "Nawafil"

	// 2. var namaVariabel = value
	var firstName = "Sentus"

	// 3. shorthand declaration
	lastName := "Agus"

	// 4. zero value declaration.
	// zero value adalah nilai default dari tipe data
	// trainingDay = ""
	var trainingDay string

	fmt.Println(
		"family name ", familyName,
		"first name ", firstName,
		"last name ", lastName,
		"training day ", trainingDay,
		"umur private ", umurPrivate,
		"umur public ", UmurPublic)
}
