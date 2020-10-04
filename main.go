package main

import (
	"github.com/kohkb/book_record/db"
	"github.com/kohkb/book_record/router"
)

func main() {
	dbConn := db.Init()
	router.Router(dbConn)
}
