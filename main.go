package main

import (
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html/v2"
	"github.com/ramnath-1998/myblog/controllers"
	"github.com/ramnath-1998/myblog/initializers"
)

func init() {
	initializers.LoadEnvVars()
}

func main() {

	engine := html.New("./views", ".html")
	app := fiber.New(fiber.Config{
		Views: engine,
	})

	app.Static("/", "./public")

	frontEndRoutes := []string{
		"/",
		"/home",
		"/about",
		"/posts",
	}

	for _, route := range frontEndRoutes {

		app.Get(route, controllers.Home)

	}

	// Start app
	app.Listen(":" + os.Getenv("PORT"))
}
