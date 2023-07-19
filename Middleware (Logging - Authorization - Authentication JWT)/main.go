package main

import (
	"myApp/db"
	"myApp/route"
)

func main() {
	db.Init()

	e := route.Init()

	e.Logger.Fatal(e.Start(":8000"))
}
