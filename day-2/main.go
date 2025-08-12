package main

import (
	interfaces "day-2/interface"
	"fmt"
)

func main() {
	// teacher1 := interfaces.Teacher{}

	// student1 := interfaces.Student{}

	// interfaces.GetHuman(teacher1)
	// interfaces.GetHuman(student1)

	sqlInstance := interfaces.Redis{}

	_, err := GetUsers(sqlInstance)
	if err != nil {
		fmt.Println("err")
	}

}

func GetUsers(db interfaces.Database) ([]interfaces.User, error) {
	return db.Find()
}
