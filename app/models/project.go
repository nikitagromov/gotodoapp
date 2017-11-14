package models

import (
	"github.com/jinzhu/gorm"
	"fmt"
)

type Project struct {
	gorm.Model
	Name string
	Tasks []Task
	OwnerID string `gorm:"ForeignKey:UserRefer"`
}


func (project Project) GetData () map[string]interface{} {
	data := make(map[string]interface{})
	data["name"] = project.Name
	data["tasks"] = project.Tasks
	data["owner_id"] = project.OwnerID
	return data
}


func CreateProject(name string, ownerId string) (project *Project) {
	project = &Project{Name:name,  OwnerID:ownerId}
	Database.Debug().Create(project)
	return
}
