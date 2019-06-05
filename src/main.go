package main

import (
	"db"
	"router"
)

func main () {
	db.OpenDB()
	defer db.SqlDB.Close()

	r := router.RegisterHandlers()
	if err := r.Run(":8000"); err != nil {
		panic(err)
	}
}

