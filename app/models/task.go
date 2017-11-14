package models

import (
	"github.com/jinzhu/gorm"
)

type Task struct {
	gorm.Model
	Name string
	ProjectID uint
	OwnerID string `gorm:"ForeignKey:UserRefer"`
}


func (task Task) GetData () map[string]interface{} {
	data := make(map[string]interface{})
	data["name"] = task.Name
	data["project_id"] = task.ProjectID
	data["owner_id"] = task.OwnerID
	return data
}

func CreateTask(name string, projectId uint, ownerId string) (task *Task) {
	task = &Task{Name:name, ProjectID:projectId, OwnerID:ownerId}
	Database.Debug().Create(task)
	return
}
