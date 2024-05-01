package Routes

import (
	"app/Controller"
	"log"

	"github.com/gofiber/fiber/v2"
)

func Task_Queue(app *fiber.App) {

	log.Println("Setting up routes for et-jenkins-notification.")
	notificationQueue := app.Group("/queue")
	jenkinsNotification := notificationQueue.Group("/")
	jenkinsNotification.Post("/et_jenkins", Controller.JenkinsNotificationQueue)
}
