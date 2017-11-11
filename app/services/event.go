package services

import "todoapp/app/models"

type EventBus struct {
	Handlers map[string]func(model *models.Model)
}

const TASK_CREATED  = "task_created"
const PROJECT_CREATED  = "project_created"


func (eventBus *EventBus) AddHandler(eventName string, handler func(model *models.Model)) {
	eventBus.Handlers[eventName] = handler
}

func (eventBus *EventBus) Dispatch(eventName string, model models.Model) {
	if eventBus.Handlers[eventName] != nil {
		handler := eventBus.Handlers[eventName]
		handler(&model)
	}
}
