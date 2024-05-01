package main

import (
	dbutil "app/Database"
	"log"

	routes "app/Routes"

	"github.com/gofiber/fiber/v2"
)

func init() {

	log.Println("Loading enviroment variables.")
	dbutil.LoadENVVar()
	dbutil.ConnectDB()
	dbutil.SyncDatabase()
}

func main() {

	log.Println("Starting the application.")
	app := fiber.New()
	routes.Task_Queue(app)
	log.Println("Application listening on port: 3000")
	app.Listen(":3000")
}
