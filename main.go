package main

import (
	"crypto/tls"
	"database/sql"
	"fmt"
	"os"

	"github.com/go-sql-driver/mysql"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html/v2"
	"github.com/gofiber/websocket/v2"
	"github.com/ramnath-1998/poorhouse/controllers"
	"github.com/ramnath-1998/poorhouse/initializers"
)

func init() {
	initializers.LoadEnvVars()

	mysql.RegisterTLSConfig("tidb", &tls.Config{
		MinVersion: tls.VersionTLS12,
		ServerName: "gateway01.eu-central-1.prod.aws.tidbcloud.com",
	})

}

func main() {

	// Running the App Engine

	engine := html.New("./views", ".html")
	app := fiber.New(fiber.Config{
		Views: engine,
	})

	// Handling Front End Routes
	frontEndRoutes := []string{
		"/",
		"/home",
		"/about",
		"/posts",
	}

	for _, route := range frontEndRoutes {

		app.Get(route, controllers.Home)

	}

	// Connecting to TiDB

	db, err := sql.Open("mysql", "2aMS5nNzrFnuQ6N.root:xtQO0UoucKUOfPOq@tcp(gateway01.eu-central-1.prod.aws.tidbcloud.com:4000)/test?tls=tidb")

	if err != nil {
		fmt.Println(err)
		return
	}

	db.Ping()

	// Getting News from API Server

	result, err := controllers.GetNewsFromApi("https://api.newscatcherapi.com/v2/search?q=all&countries=IN&lang=en&topic=economics&page_size=500")

	if err != nil {
		fmt.Println(err)
		return
	}

	articles := result.Articles
	for _, article := range articles {
		fmt.Println(article.Title)
	}

	// Handling Backend Routes
	// websocket route
	app.Static("/", "./public")
	app.Get("/ws", websocket.New(controllers.HandleWebSocket))

	// Start app
	app.Listen(":" + os.Getenv("PORT"))
	fmt.Print("Its Running!")

}
