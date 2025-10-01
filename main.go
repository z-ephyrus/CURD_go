package main

import (
	"CURD/internals/db"
	"fmt"
	"os"
)

func main() {


	fmt.Println("DB connection testing (for pg)")

	dsn := os.Getenv("DSN")
	conn := db.Connect(dsn)

	

	defer conn.Close()
}