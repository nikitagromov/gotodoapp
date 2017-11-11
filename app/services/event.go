package services

import (
	"todoapp/app/models"
	"fmt"
)

type EventBus struct {
	channels map[string]chan *models.Model
}

const TASK_CREATED  = "task_created"
const PROJECT_CREATED  = "project_created"


func (eventBus *EventBus) AddHandler(eventName string, handler func(model *models.Model)) {
	channel := eventBus.channels[eventName]
	go handlerWrapper(handler, channel)
}


func handlerWrapper(handler func(model *models.Model), ch chan *models.Model) {
	for {
		model := <- ch
		handler(model)
	}
}

func (eventBus *EventBus) Dispatch(eventName string, model models.Model) {
	if eventBus.channels[eventName] != nil {
		eventBus.channels[eventName] <- &model
	}
}


func (eventBus *EventBus) Init () {
	eventBus.channels = make(map[string]chan *models.Model)

	for _, event := range []string {TASK_CREATED, PROJECT_CREATED} {
		eventBus.channels[event] = make(chan *models.Model)
	}

}