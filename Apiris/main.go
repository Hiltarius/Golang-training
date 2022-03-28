package main


import (
	"github.com/kataras/iris/v12"
	"fmt"
	"github.com/iris-contrib/middleware/cors"
)

func main() {
	app := newApp()

    app.Run(iris.Addr(":8080"), iris.WithoutServerError(iris.ErrServerClosed))
}


func newApp() *iris.Application {
	app := iris.New()

	crs := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedHeaders:   []string{"*"},
		AllowCredentials: true,
	})

    app.UseRouter(crs)

	app.Get("/", func(ctx iris.Context){
		ctx.Text("Hello World")

	})
	app.Get("{name:string}", func(ctx iris.Context) {
		ctx.Text(fmt.Sprintf("Hello %s", ctx.Params().Get("name")))

	})
	app.Get("/lol/{ok:string}", func(ctx iris.Context) {
		ctx.Text(fmt.Sprintf("Hello %s", ctx.Params().Get("ok")))

	})
	return app
}
 

