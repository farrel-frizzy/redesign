package main

import (
	"fmt"
	"redesign/app"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db := app.NewDB()

	if db != nil {
		fmt.Println("Connected to the database")
	} else {
		fmt.Println("Failed to connect to the database")
	}
}
