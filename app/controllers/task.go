package controllers

import (
	"todoapp/app"
	"todoapp/app/models"
	"github.com/revel/revel"
	"encoding/json"
	"fmt"
	"net/url"
	"github.com/jinzhu/gorm"
	"strconv"
)

type TaskJSON struct {
	Name string `json:"name"`
	ProjectID uint `json:"projectID"`
}

type (TaskController struct {
	*revel.Controller

})


func (c TaskController) GetTasksCollectionView() revel.Result {
	tasks := []models.Task{}
	query := getItemsCollectionQuery(c.Params.Query)
	query = addProjectId(c.Params.Query, query)
	query = processQParam(c.Params.Query, query)
	query.Find(&tasks)
	return c.RenderJSON(tasks)
}

func (c TaskController) GetTaskById() revel.Result {
	task := models.Task{}
	app.Database.Debug().Where("id = ?", c.Params.Get("id")).First(&task)
	fmt.Println(task)
	return c.RenderJSON(task)
}

func (c TaskController) AddTask() revel.Result {
	var payload TaskJSON
	data, _ := getBody(c.Request)
	json.Unmarshal(data, &payload)
	task := models.Task{Name: payload.Name, ProjectID: payload.ProjectID}
	app.Database.Debug().Create(&task)
	return c.RenderJSON(task)
}

func addProjectId(c url.Values, db *gorm.DB)  *gorm.DB {
	projectId, error := strconv.Atoi(c.Get("project_id"))
	query := db

	if error == nil {
		query = query.Where("project_id = ?", projectId)
	}

	return query
}

