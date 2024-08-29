package main

import (
	"fmt"
	"log"

	app "github.com/nguyenluan2001/golang-authenticate/pkg/app"
	"github.com/nguyenluan2001/golang-authenticate/pkg/database"
	"github.com/valyala/fasthttp"
)

type User struct {
	email string
}

func main() {
	user := &User{
		email: "luannguyen",
	}
	fmt.Println(user)
	fmt.Println("Server start")
	// r := router.New()
	// r.GET("/", Index)
	// r.GET("/hello/{name}", Hello)
	db := &database.Postgres{}
	db.Connect()
	db.Migrate()

	app := &app.App{
		DB: db.DB,
	}
	handler := app.InitRouter().Handler

	log.Fatal(fasthttp.ListenAndServe(":8081", handler))
}
