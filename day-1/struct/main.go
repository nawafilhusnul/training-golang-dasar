package main

import "fmt"

// STRUCT
// struct adalah tipe data yang memiliki nama properti dan tipe data
// struct bisa di deklarasi dengan nama
type user struct {
	firstName string
	lastName  string
	age       int
	money     float64
	hobbies   []string
}

type employee struct {
	user    user
	company company
}

type company struct {
	name    string
	address string
	unitNo  int
	year    int
}

// struct method adalah function yang digunakan untuk mengakses properti struct
// struct method bisa menggunakan pointer atau value
func (u *user) GetFullName() string {
	return u.firstName + " " + u.lastName
}

// struct method dengan pointer
func (u *user) ChangeFirstName(s string) {
	u.firstName = s
}

// struct method dengan value
func (u user) ChangeLastName(s string) {
	u.lastName = s
}

func main() {

	// cara deklarasi struct
	user1 := user{
		firstName: "Agus",
		lastName:  "Budi",
		age:       21,
		money:     50000,
		hobbies:   []string{"football", "rugby"},
	}

	fmt.Println(user1.GetFullName())

	// nama user1 akan berubah karena menggunakan pointer
	user1.ChangeFirstName("Budi")
	fmt.Println(user1.GetFullName())

	// nama user1 tidak akan berubah karena tidak menggunakan pointer
	user1.ChangeLastName("Widodo")
	fmt.Println(user1.GetFullName())

	// deklarasi embedded struct
	employee1 := employee{
		user: user{
			firstName: "Agus",
			lastName:  "Budi",
			age:       21,
			money:     50000,
			hobbies:   []string{"football", "rugby"},
		},
		company: company{
			name:    "Agus",
			address: "Budi",
			unitNo:  21,
			year:    2021,
		},
	}

	fmt.Println(employee1.user.GetFullName())

	// deklarasi embedded struct dengan cara lain
	employee2 := employee{}
	userEmployee2 := user{
		firstName: "Agus",
		lastName:  "Budi",
		age:       21,
		money:     50000,
		hobbies:   []string{"football", "rugby"},
	}

	companyEmployee2 := company{
		name:    "Coding Studio",
		address: "Budiluhur",
		unitNo:  21,
		year:    2021,
	}

	employee2.user = userEmployee2
	employee2.company = companyEmployee2

	fmt.Println(employee2.user.GetFullName(), employee2.company.name)
}
