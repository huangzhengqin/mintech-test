package main

import (
	"db"
	"fmt"
	"router"
)

func main () {
	db.OpenDB()
	defer db.SqlDB.Close()

	r := router.RegisterHandlers()
	err := r.Run(":8000")
	if err != nil {
		fmt.Println(err)
	}
}

