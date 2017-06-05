package controllers

import (
	"todoapp/app"
	"todoapp/app/models"
	"github.com/revel/revel"
	"io/ioutil"
	"encoding/json"
)

type ProjectJSON struct {
	Name string `json:"name"`
	Tasks []uint `json:"tasks"`
}

type (ProjectController struct {
	*revel.Controller
})

func (c ProjectController)  GetBody() ([]byte, error) {
	return ioutil.ReadAll(c.Request.Body)
}

func (c ProjectController) GetQueryParams(param string) string {
	return c.Params.Query.Get(param)
}

func (c ProjectController) GetParams(param string) string {
	return c.Params.Get(param)
}


func (c ProjectController) GetProjectsCollectionView() revel.Result {
	projects := []models.Project{}
	query := getItemsCollectionQuery(c.Params.Query)
	query = processQParam(c.Params.Query, query)
	query.Find(&projects)
	return c.RenderJSON(projects)
}

func (c ProjectController) GetProjectById() revel.Result {
	project := models.Project{}
	app.Database.Preload("Tasks").First(&project, "id = ?", c.Params.Get("id"))
	return c.RenderJSON(project)
}

func (c ProjectController) AddProject() revel.Result {
	var payload ProjectJSON
	data, _ := c.GetBody()
	json.Unmarshal(data, &payload)
	project := models.Project{Name: payload.Name}
	app.Database.Debug().Create(&project)
	return c.RenderJSON(project)
}

