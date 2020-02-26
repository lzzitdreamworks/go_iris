package main

import (
	"github.com/kataras/iris/v12"

	"github.com/kataras/iris/v12/middleware/logger"
	"github.com/kataras/iris/v12/middleware/recover"
)

type Person struct {
	Username string `json:"username"`
	Pwd string `json:"pwd"`
}


func main() {
	app := iris.New()
	app.Logger().SetLevel("debug")
	// Optionally, add two builtin handlers
	// that can recover from any http-relative panics
	// and log the requests to the terminal.
	app.Use(recover.New())
	app.Use(logger.New())

	// Method:   GET
	// Resource: http://localhost:8080
	app.Handle("GET", "/", func(ctx iris.Context) {
		ctx.HTML("<h1>Go Web Iris, Welcome</h1>")
	})

	// same as app.Handle("GET", "/ping", [...])
	// Method:   GET
	// Resource: http://localhost:8080/ping
	app.Get("/ping", func(ctx iris.Context) {
		path := ctx.Path
		app.Logger().Info(path)
		username := ctx.URLParam("username")
		app.Logger().Info(username)
		password := ctx.URLParam("pwd")
		app.Logger().Info(password)
	//	ctx.WriteString(string(path))
		ctx.HTML("<H1>" + username + "," + password + "</h1>")
	})

	// Method:   GET
	// Resource: http://localhost:8080/hello
	app.Get("/hello", func(ctx iris.Context) {
		ctx.JSON(iris.Map{"message": "Hello Iris!"})
	})


	app.Post("/login", func(ctx iris.Context) {
		path := ctx.Path
		app.Logger().Info(path)
		username := ctx.PostValue("username")
		app.Logger().Info(username)
		password := ctx.PostValue("pwd")
		app.Logger().Info(password)
	//	ctx.WriteString(string(path))
		ctx.HTML("<H1>" + username + "," + password + "</h1>")
	})

	app.Post("/postjson", func(ctx iris.Context) {
		path := ctx.Path
		app.Logger().Info(path)
		var person Person
		if err := ctx.ReadJSON(&person); err != nil {
			panic(err.Error())
		}
		
	//	ctx.WriteString(string(path))
		ctx.Writef("Received", person)
	})

	// http://localhost:8080
	// http://localhost:8080/ping
	// http://localhost:8080/hello
	app.Run(iris.Addr(":8080"), iris.WithoutServerError(iris.ErrServerClosed))
}
