package controllers


import (
	"todoapp/app/models"
	"github.com/revel/revel"
	"encoding/json"
	"fmt"
)

type UserJSON struct {
	Email string `json:"email"`
	Password string `json:"password"`
}

type FilteredUserJSON struct {
	Email string `json:"email"`
}

type (UserController struct {
	*revel.Controller

})


func (c UserController) GetUsers() revel.Result {
	users := []models.User{}
	query := getItemsCollectionQuery(c.Params.Query)
	query = processQParam(c.Params.Query, query)
	query.Find(&users)
	filteredUsers := []FilteredUserJSON{}

	for i := range users {
		filteredUsers = append(filteredUsers, FilteredUserJSON{Email:users[i].Email})
	}

	return c.RenderJSON(filteredUsers)
}

func (c UserController) GetUserByEmail() revel.Result {
	user := models.User{}
	email :=  c.Params.Get("email")
	models.Database.Debug().Where("email = ?", email).First(&user)
	return c.RenderJSON(&FilteredUserJSON{Email: user.Email})
}

func (c UserController) AddUser() revel.Result {
	var payload UserJSON
	data, _ := getBody(c.Request)
	json.Unmarshal(data, &payload)
	user := models.User{Email: payload.Email, Password: models.EncodePassword(payload.Password)}
	err := models.Database.Debug().Create(&user).Error
	if err != nil {
		fmt.Println(err.Error())
		if err.Error() == "pq: duplicate key value violates unique constraint \"users_pkey\"" {
			c.Response.Status = 409
			return c.RenderError(err)
		} else {
			return c.RenderError(err)
		}

	}

	return c.RenderJSON(&FilteredUserJSON{Email: user.Email})
}
