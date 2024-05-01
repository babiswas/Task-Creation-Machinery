package Controller

import (
	"log"
	"os"

	machinery "github.com/RichardKnop/machinery/v1"
	"github.com/RichardKnop/machinery/v1/config"
	"github.com/RichardKnop/machinery/v1/tasks"
	"github.com/gofiber/fiber/v2"
)

func JenkinsNotificationQueue(ctx *fiber.Ctx) error {

	var cnf = config.Config{
		Broker:        os.Getenv("REDIS_SERVER_URL"),
		ResultBackend: os.Getenv("REDIS_SERVER_URL"),
	}

	var body struct {
		JobName        string `json:"jenkins_job_name"`
		ProjectName    string `json:"project_name"`
		FeatureName    string `json:"feature_name"`
		Status         string `json:"status"`
		BuildNumber    string `json:"build_number"`
		ChildJob       string `json:"child_job"`
		JobParamsKey   string `json:"job_params_key"`
		JobParamsValue string `json:"job_params_value"`
	}

	err := ctx.BodyParser(&body)
	log.Println(body)

	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"success": false, "message": "InvalidData"})
	}

	server, err := machinery.NewServer(&cnf)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"success": false, "message": "Failed to queue the task."})
	}

	log.Println("Task:", body)

	sayTask := tasks.Signature{
		Name: "process_et_jenkins_notification",
		Args: []tasks.Arg{
			{
				Name:  "JobName",
				Type:  "string",
				Value: body.JobName,
			},
			{
				Name:  "ProjectName",
				Type:  "string",
				Value: body.ProjectName,
			},
			{
				Name:  "FeatureName",
				Type:  "string",
				Value: body.FeatureName,
			},
			{
				Name:  "BuildNumber",
				Type:  "string",
				Value: body.BuildNumber,
			},
			{
				Name:  "Status",
				Type:  "string",
				Value: body.Status,
			},
			{
				Name:  "childJob",
				Type:  "string",
				Value: body.ChildJob,
			},
			{
				Name:  "jobParamKeys",
				Type:  "string",
				Value: body.JobParamsKey,
			},
			{
				Name:  "jobParamVals",
				Type:  "string",
				Value: body.JobParamsValue,
			},
		},
	}

	server.SendTask(&sayTask)
	return ctx.Status(fiber.StatusCreated).JSON(fiber.Map{"sucess": true, "message": "Task Queued."})
}
