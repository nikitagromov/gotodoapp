package controllers

import (
	"todoapp/app/models"
	"github.com/revel/revel"
	"encoding/json"
	"fmt"
	"net/url"
	"github.com/jinzhu/gorm"
	"strconv"
	"todoapp/app"
	"todoapp/app/services"
)

type TaskJSON struct {
	Name string `json:"name"`
	ProjectID uint `json:"projectID"`
	OwnerID string `json:"ownerID"`
}

type (TaskController struct {
	*revel.Controller

})


func (c TaskController) GetTasksCollectionView() revel.Result {
	tasks := []models.Task{}
	query := getItemsCollectionQuery(c.Params.Query)
	query = filterByProjectId(c.Params.Query, query)
	query = processQParam(c.Params.Query, query)
	query.Find(&tasks)
	return c.RenderJSON(tasks)
}

func (c TaskController) GetTaskById() revel.Result {
	task := models.Task{}
	models.Database.Debug().Where("id = ?", c.Params.Get("id")).First(&task)
	fmt.Println(task)
	return c.RenderJSON(task)
}

func (c TaskController) AddTask() revel.Result {
	var payload TaskJSON
	data, _ := getBody(c.Request)
	json.Unmarshal(data, &payload)
	task := models.CreateTask(payload.Name, payload.ProjectID, payload.OwnerID)
	app.EventBus.Dispatch(services.TASK_CREATED, task)
	return c.RenderJSON(task)

}

func filterByProjectId(c url.Values, db *gorm.DB)  *gorm.DB {
	projectId, err := strconv.Atoi(c.Get("project_id"))
	query := db

	if err == nil {
		query = query.Where("project_id = ?", projectId)
	}

	return query
}

