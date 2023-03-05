package main

import (
	"fmt"
	"github.com/dre-zouhair/mailer/internal/db"
)

func main() {
	var connection, err = db.Connect()
	if err != nil {
		fmt.Println("unable to init connection with redis db :", err)
		return
	} else {
		fmt.Println("Connected ", connection.GetDB().ClientID())
	}

	defer closeConnect(connection)
}

func closeConnect(connection *db.Database) {
	err := connection.Close()
	if err != nil {
		fmt.Println("unable to close redis connection")
		return
	} else {
		fmt.Println("Disconnected")
	}
}
