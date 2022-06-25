package main

import (
	"github.com/Haffif/web-echo/db"
	"github.com/Haffif/web-echo/routes"
)

func main() {
	db.Init()

	e := routes.Init()
	e.Logger.Fatal(e.Start(":1323"))
}
