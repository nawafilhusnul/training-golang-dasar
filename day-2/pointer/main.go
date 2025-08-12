package main

import "fmt"

func main() {
	var age int
	fmt.Println(age)
	fmt.Println(&age)

	user1 := User{}
	user1.SetName("Sentus")

	fmt.Println("harapannya sentus", user1.Name)
}

type User struct {
	Name string
}

func (u *User) SetName(name string) {
	u.Name = name
}
