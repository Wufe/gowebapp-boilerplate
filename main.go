package main

import (
	"fmt"

	"github.com/wufe/boilerplateprj/infrastructure"
	"github.com/wufe/boilerplateprj/presentation"
	"github.com/wufe/boilerplateprj/services"

	"github.com/kataras/iris"
	"github.com/kataras/iris/middleware/logger"
	"github.com/kataras/iris/middleware/recover"

	"go.uber.org/dig"
)

func registerDatabase(diContainer *dig.Container) {
	database := infrastructure.NewDatabase()
	database.Connect()
	database.Automigrate()
	database.Seed()

	diContainer.Provide(func() infrastructure.DatabaseAccessor {
		return database
	})
}

func registerServices(diContainer *dig.Container) {
	servicesMap := make(map[string]interface{})

	// Create an entry here to register a service via IOC
	servicesMap["Home"] = services.NewHomeService

	for key, service := range servicesMap {
		err := diContainer.Provide(service)
		if err != nil {
			fmt.Printf("Could not resolve dependencies for service %s\n", key)
			fmt.Println(err.Error())
		}
	}
}

// Starts up HTTP handlers by invocation
// Should be called after all DI registrations
func startupHTTPHandlers(app *iris.Application, diContainer *dig.Container) {
	handlersMap := make(map[string]interface{})

	// Add a new controller here
	handlersMap["/home"] = presentation.GetHomeHandlers(app.Party("/home"))

	for key, handler := range handlersMap {
		err := diContainer.Invoke(handler)
		if err != nil {
			fmt.Printf("Could not resolve dependencies for handler %s\n", key)
			fmt.Println(err.Error())
		}
	}

}

func main() {
	app := iris.New()

	app.Logger().SetLevel("error")
	app.Use(recover.New())
	app.Use(logger.New())

	// app.RegisterView(iris.HTML("./dist/static", ".html"))
	app.Get("/", func(ctx iris.Context) {
		ctx.ServeFile("./dist/static/index.html", true)
		// ctx.View("index.html")
	})

	app.StaticWeb("/static", "./dist/static")

	diContainer := dig.New()
	registerDatabase(diContainer)
	registerServices(diContainer)

	startupHTTPHandlers(app, diContainer)

	port := 2345
	fmt.Printf("Starting server on port %d..\n", port)
	// app.OnErrorCode(iris.StatusInternalServerError, func(ctx iris.Context) {
	// 	ctx.HTML("InternalServerError: <b>" + ctx.Values().GetString("message") + "</b>")
	// })
	app.Run(iris.Addr(fmt.Sprintf(":%d", port)), iris.WithoutServerError(iris.ErrServerClosed), iris.WithoutStartupLog)
	// app.Run(iris.Addr(fmt.Sprintf(":%d", port)))

}
