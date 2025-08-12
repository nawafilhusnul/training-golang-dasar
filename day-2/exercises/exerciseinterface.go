package exercises

type Shape interface {
	Luas() float64
	Keliling() float64
}

// Luas Lingkaran : radius^2 * 3,14
// Keliling Lingkaran : 2*radius*3,14
type Lingkaran struct {
	// Tambahin yg hilang
	radius float64
}

// luas sisixsisi
// keliling sisix4
type Persegi struct {
	panjangSisi float64
}
