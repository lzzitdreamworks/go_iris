package main

import (
	"fmt"
	"github.com/go-xorm/xorm"
	"github.com/kataras/iris/v12"
	_ "github.com/go-sql-driver/mysql"
	"github.com/kataras/iris/v12/middleware/logger"
	"github.com/kataras/iris/v12/middleware/recover"
	"log"
)

const DriverName = "mysql"
// "数据库账号:数据库密码@tcp(数据库Ip:端口)/数据库名称?parseTime=true"
const MasterDataSourceName = "root:@tcp(127.0.0.1:3306)/otott_slave?charset=utf8"

var engine *xorm.Engine

func execute () {
    sql := "INSERT INTO `admin`(`id`, `username`, `password`, `true_name`, `email`, `mobile`, `remark`, `create_admin_id`, `last_login_ip`, `last_login_time`, `login_count`, `last_logout_time`, `create_time`, `update_time`, `status`) VALUES (29, 'sz_admin', '$2y$13$nHT706UX/3qDO2FunpzcquY0YxLDw3X9buoXoVTITA3dLXT1bzFxq', '深圳admin', NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL, '2019-12-20 14:48:49', '2019-12-20 14:48:49', 1);"
    affected, err := engine.Exec(sql)
    if err != nil {
        log.Fatal("execute error", err)

    } else {
        id, _ := affected.LastInsertId()
        rows, _ := affected.RowsAffected()
        fmt.Println("execute id=", id, ",rows=", rows)
    }
}

// 连接到数据库
func newEngin() *xorm.Engine {
	engine, err := xorm.NewEngine(DriverName, MasterDataSourceName)
	if err != nil {
		log.Fatal(newEngin, err)
		return nil
	}
	// Debug模式，打印全部SQL语句，帮助对比
	engine.ShowSQL(true)
	return engine
}

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

	engine = newEngin()
	execute() // 测试插入一条语句

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
