package main

import "fmt"

func main() {

	// INT (integer)
	// int adalah tipe data yang memiliki nilai negatif dan positif
	// int8, int16, int32, int64 adalah tipe data yang memiliki 8, 16, 32, 64 bit
	var numInt int     // zero value adalah 0
	var numInt8 int8   // zero value adalah 0
	var numInt16 int16 // zero value adalah 0
	var numInt32 int32 // zero value adalah 0
	var numInt64 int64 // zero value adalah 0

	fmt.Println("numInt", numInt, "numInt8", numInt8, "numInt16", numInt16, "numInt32", numInt32, "numInt64", numInt64)

	// UINT (unsigned integer)
	// uint adalah tipe data yang hanya memiliki nilai positif
	// angka positifnya yang ditampung akan lebih besar
	// uint8, uint16, uint32, uint64 adalah tipe data yang memiliki 8, 16, 32, 64 bit
	var numUint uint     // zero value adalah 0
	var numUint8 uint8   // zero value adalah 0
	var numUint16 uint16 // zero value adalah 0
	var numUint32 uint32 // zero value adalah 0
	var numUint64 uint64 // zero value adalah 0

	fmt.Println("numUint", numUint, "numUint8", numUint8, "numUint16", numUint16, "numUint32", numUint32, "numUint64", numUint64)

	// FLOAT
	// float adalah tipe data yang memiliki nilai negatif dan positif
	// float32 adalah tipe data yang memiliki 32 bit
	// float64 adalah tipe data yang memiliki 64 bit
	// biasanya digunakan untuk decimal
	var myFloat32 float32

	var myFloat64 float64

	// cara deklarasi tipe data int, float, uint
	var myWifeAge int = 20            // int
	myAge := 21                       // int
	myMoney := 500000.00              // float64
	myNumUint1 := uint(10)            // uint
	var myFloatMoney float64 = 500000 // float64 ; // deklarasi jika sudah ada tipe data bisa pake koma ataupun tidak

	fmt.Println("myFloat32", myFloat32, "myFloat64", myFloat64, "myWifeAge", myWifeAge, "myAge", myAge, "myMoney", myMoney, "myNumUint1", myNumUint1, "myFloatMoney", myFloatMoney)

	// // BOOLEAN
	// boolean adalah tipe data yang hanya memiliki 2 nilai yaitu true atau false
	var isMarried bool // zero value adalah false
	isMarried = true   // assign value

	// boolean bisa di deklarasi langsung
	isMarried2 := true

	fmt.Println("isMarried", isMarried, "isMarried2", isMarried2)

	// INTERFACE KOSONG
	// interface kosong adalah tipe data yang bisa menampung semua tipe data
	var apapun interface{}
	var apapun2 any
	apapun = 21
	apapun = "Name"
	apapun = 21

	fmt.Println("apapun", apapun, "apapun2", apapun2)

	// STRING
	// string adalah tipe data teks
	// string bisa di deklarasi langsung dengan tanda petik dua, namun hanya satu line
	var myName string = "Husnul"
	myName2 := "Agus"

	// string bisa di deklarasi langsung dengan backtick juga dan bisa multi line
	fileDescription := `ini adalah file description
		ini adalah file description
		ini adalah file description
	`

	fmt.Println("myName", myName, "myName2", myName2, "fileDescription", fileDescription)

	// KONSTANTA
	// konstanta adalah tipe data yang tidak bisa diubah nilainya
	// biasanya digunakan untuk nilai yang tidak akan berubah
	// cara deklarasi konstanta sama dengan yang lainnya cuma var diganti menjadi const
	// constant tidak bisa dideklarasi dengan metode shorthand
	const companyName = "BANK PAPUA"

	// POINTER
	// pointer adalah alamat memori dari suatu variable
	// & adalah operator untuk mendapatkan alamat memori
	// * adalah operator untuk mengakses alamat memori
	// zero value dari pointer adalah nil

	// deklarasi sebuah variable
	a := "saya"

	// deklarasi sebuah variable pointer
	b := &a

	// mengubah nilai dari variable pointer
	// * adalah operator untuk mengakses alamat memori
	*b = "kamu"

	// deklarasi pointer kosong
	var stringA *string // zero value adalah nil

	fmt.Println("a", a, "b", *b, "stringA", stringA)

	// EXERCISE 1:
	//
	//  "Saya {{nama}} umur saya {{umur}}. Uang saya di kantong ada {{uang}}."
	namaSaya := "Husnul"
	umurSaya := 21
	uangSaya := 500000.00

	// fmt.Sprintf adalah format string. String ini akan di format sesuai dengan format yang diberikan
	// %s adalah string
	// %d adalah int
	// %.2f adalah float dengan 2 decimal
	// %t adalah boolean
	// %v adalah tipe data apapun
	// resources : https://yourbasic.org/golang/fmt-printf-reference-cheat-sheet/
	res := fmt.Sprintf("Saya %s, umur %d tahun. Uang saya %.2f", namaSaya, umurSaya, uangSaya)
	fmt.Println(res)

}
