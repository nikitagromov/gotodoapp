package controllers

import (
	"todoapp/app/models"
	"github.com/revel/revel"
	"encoding/json"
	"todoapp/app"
	"todoapp/app/services"
)

type ProjectJSON struct {
	Name string `json:"name"`
	Tasks []uint `json:"tasks"`
}

type (ProjectController struct {
	*revel.Controller
})

func (c ProjectController) GetProjectsCollectionView() revel.Result {
	projects := []models.Project{}
	query := getItemsCollectionQuery(c.Params.Query)
	query = processQParam(c.Params.Query, query)
	query = query.Preload("Tasks")
	query.Find(&projects)
	return c.RenderJSON(projects)
}

func (c ProjectController) GetProjectById() revel.Result {
	project := models.Project{}
	models.Database.Preload("Tasks").First(&project, "id = ?", c.Params.Get("id"))
	return c.RenderJSON(project)
}

func (c ProjectController) AddProject() revel.Result {
	var payload ProjectJSON
	data, _ := getBody(c.Request)
	json.Unmarshal(data, &payload)
	project := models.Project{Name: payload.Name}
	models.Database.Debug().Create(&project)
	app.EventBus.Dispatch(services.PROJECT_CREATED, project)
	return c.RenderJSON(project)
}

