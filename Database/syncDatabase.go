package Database

import (
	"app/Model"
	"log"
)

func SyncDatabase() {
	log.Println("Performing Database Migration.")
	DB.AutoMigrate(&Model.JenkinsJobStatus{})
}
