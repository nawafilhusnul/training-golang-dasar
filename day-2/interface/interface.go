package interfaces

import "fmt"

type Manusia interface {
	Walk()
	Speak()
	GetMarried()
}

type Teacher struct{}

func (t Teacher) Walk() {
	fmt.Println("Teacher Walking")
}
func (t Teacher) Speak() {
	fmt.Println("Teacher Speaks")
}
func (t Teacher) GetMarried() {
	fmt.Println("Teacher Get Married")
}

type Student struct{}

func (t Student) Walk() {
	fmt.Println("Student Walking")
}
func (t Student) Speak() {
	fmt.Println("Student Speaks")
}
func (t Student) GetMarried() {
	fmt.Println("Student Get Married")
}

func GetHuman(h Manusia) {
	h.Speak()
	h.Walk()
	h.GetMarried()
}

type User struct{}

type Database interface {
	Insert(u User) error
	Find() ([]User, error)
}

type MySQL struct{}

func (sql MySQL) Insert(u User) error {
	_ = "INSERT INTO users VALUES...."
	return nil
}

func (sql MySQL) Find() ([]User, error) {
	_ = "SELECT * ......"

	return []User{}, nil
}

type Redis struct{}

func (rd Redis) Insert(u User) error {
	_ = "Query redis"
	return nil
}

func (rd Redis) Find() ([]User, error) {
	_ = "QueryRedisf"

	return []User{}, nil
}
