package main

import (
	"os"

	"github.com/djarum76-bot/crud_post/db"
	"github.com/djarum76-bot/crud_post/routes"
)

func main() {
	db.Init()

	e := routes.Init()

	port := os.Getenv("PORT")

	e.Logger.Fatal(e.Start(":" + port))
}
