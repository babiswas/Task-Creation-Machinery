package Task

import (
	"app/Database"
	model "app/Model"
	"log"
)

func Process_ET_Jenkins_Notification(JobName, ProjectName, FeatureName, BuildNumber string, Status string, childJob string, jobParamKeys string, jobParamVals string) (string, error) {
	var job_status bool
	if Status == "true" {
		job_status = true
	} else {
		job_status = false
	}
	jenkins_job := model.JenkinsJobStatus{JobName: JobName, ProjectName: ProjectName, FeatureName: FeatureName, Status: job_status, BuildNumber: BuildNumber, ChildJob: childJob, JobParamsKey: jobParamKeys, JobParamsValue: jobParamVals}
	result := Database.DB.Create(&jenkins_job)
	if result.Error != nil {
		log.Println("Error occured :", result.Error)
		return "FAILURE", result.Error
	}
	return "SUCCESS", nil
}
