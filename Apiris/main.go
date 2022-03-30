package main


import (
	"github.com/kataras/iris/v12"
	"fmt"
	"github.com/iris-contrib/middleware/cors"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
    "Apiris/model"

)


func main() {
	// Create the databse connection
	db, err := gorm.Open(
	"postgres",
	"host=localhost port=5432 user=postgres dbname=postgres password=mysecretpassword sslmode=disable",
	)
	// End the program with an error if it could not connect to the database
	if err != nil {
	panic("could not connect to database")
	}

	// Create the default database schema and a table for the orders
	db.AutoMigrate(&model.Order{})

	// Closes the database connection when the program ends
	defer func() {
	_ = db.Close()
	}()


	app := newApp(db)

    app.Run(iris.Addr(":8080"), iris.WithoutServerError(iris.ErrServerClosed))
}


func newApp(db *gorm.DB) *iris.Application {
	app := iris.New()

	crs := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedHeaders:   []string{"*"},
		AllowCredentials: true,
	})

    app.UseRouter(crs)

	// Define the slice for the result
	var orders []model.Order

	// Endpoint to perform the database request
	app.Get("/orders", func(ctx iris.Context) {
		db.Find(&orders)
		// Return the result as JSON
		ctx.JSON(orders)
	})


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
