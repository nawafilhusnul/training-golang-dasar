package main

import (
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"

	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "YdlFHWVsoUQaEwJSenU2"
	dbname   = "postgres"
)

func main() {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		fmt.Println("ERROR ")
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	usegw := "tes"
	pswd := "tes213!"
	sqlRow := db.QueryRow(
		"SELECT Id, aplikasi, ip, pswd, sts FROM usergateway WHERE usergw = $1 AND pswd = $2",
		usegw, pswd,
	)

	var usergateway UserGateway
	err = sqlRow.Scan(
		&usergateway.Id,
		&usergateway.Aplikasi,
		&usergateway.IP,
		&usergateway.Pswd,
		&usergateway.Sts,
	)

	if usergateway.Sts == "0" {
		fmt.Println("USER ANDA ACTIVE")
	}

	usergateway.Age = 10
	if err != nil {
		switch {
		case errors.Is(err, sql.ErrNoRows):
			fmt.Println("DATA TIDAK DITEMUKAN!!!")
			return
		}

		fmt.Println("FAILED TO SCAN: ", err)
		return
	}

	BuildResponseAPI("rC205", "Success to get user gateway", usergateway)
	fmt.Println("Successfully connected!")
}

type UserGateway struct {
	Id       int    `json:"userId" redis:"id"`
	Aplikasi string `json:"userApp"`
	IP       string `json:"userIP"`
	Pswd     string `json:"-"`
	Age      int    `json:"age,omitempty"`
	Sts      string `json:"-"`
}

type ResponseAPI struct {
	RCode   string `json:"rcode"`
	Message string `json:"message"`
	Data    any    `json:"data"`
}

func BuildResponseAPI(rcode, message string, data any) {
	resp := ResponseAPI{
		RCode:   rcode,
		Message: message,
		Data:    data,
	}

	b, err := json.Marshal(resp)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(string(b))
}
