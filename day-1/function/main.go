package main

import (
	"errors"
	"fmt"
)

func main() {
	// cara memanggil function
	simpleFunction()
	functionWithOneArgument("Agus")

	agusName := functionWithOneArgumentAndReturnValue("Agus")
	fmt.Println("agusName", agusName)

	functionWithMultipleArguments("Agus", 21)

	name, age := functionWithMultipleArgumentsAndReturnValue("Agus", 21)
	fmt.Println("name", name, "age", age)

	// ketika variable sudah dideklarasi, maka tidak perlu di deklarasi lagi
	name, age = functionWithMultipleReturnValue()
	fmt.Println("name", name, "age", age)

	// _ adalah variable yang tidak akan digunakan
	name, _ = functionWithNamedReturnValue()
	fmt.Println("name", name)

	functionWithVariadicArguments(1, 2, 3, 4, 5)

	// function dengan error
	err := functionWithError()
	if err != nil {
		fmt.Println("error", err)
	}

	// anonymous function
	func() {
		fmt.Println("anonymous function")
	}()

	umur := 21
	// anonymous function bisa digunakan untuk mengganti ternary operator
	isMarried := func(u int) bool {
		return u > 18
	}(umur)

	fmt.Println("isMarried", isMarried)

}

// function adalah block kode yang bisa dijalankan
// function yang tidak ada arguments dan tidak ada return value
func simpleFunction() {
	// tidak perlu return value
}

// function dengan satu argument
func functionWithOneArgument(name string) {
	// perlu return value
}

// function dengan satu argument dan return value
func functionWithOneArgumentAndReturnValue(name string) string {
	return name
}

// function dengan multiple arguments
func functionWithMultipleArguments(name string, age int) {

}

// function dengan multiple arguments dan return value
func functionWithMultipleArgumentsAndReturnValue(name string, age int) (string, int) {
	return name, age
}

// function dengan multiple return value
func functionWithMultipleReturnValue() (string, int) {
	return "name", 21
}

// function dengan named return value
func functionWithNamedReturnValue() (name string, age int) {
	name = "name"
	age = 21
	return
}

// function dengan variadic arguments
// variadic arguments adalah arguments yang bisa di deklarasi lebih dari 1
// variadic arguments harus ditaruh di akhir arguments
func functionWithVariadicArguments(args ...int) {
	for _, arg := range args {
		fmt.Println(arg)
	}
}

// function dengan error
func functionWithError() error {
	return errors.New("error")
}
