package main

import (
	"db"
	"router"
)

func main () {
	d := db.OpenDB()
	defer d.Close()

	if err := router.NewService(d).New().Run(":8000"); err != nil {
		panic(err)
	}
}

