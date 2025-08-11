package calculator

import "errors"

// Calculate adalah function yang digunakan untuk menghitung
func Calculate(a, b int, operators ...string) (int, error) {
	// jika tidak ada operator, maka akan mengembalikan hasil penjumlahan
	if len(operators) == 0 {
		return a + b, nil
	}

	// jika operator ditemukan, maka akan mengembalikan hasil sesuai operator
	switch operators[0] {
	case "+":
		return a + b, nil
	case "-":
		return a - b, nil
	case "*":
		return a * b, nil
	case "/":
		return a / b, nil
	default:
		return 0, errors.New("operator tidak dikenal")
	}
}
