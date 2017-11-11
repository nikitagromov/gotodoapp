package services

import (
	"fmt"
	"todoapp/app/models"
)

type Message struct {
	Addressee string
	Body map[string] interface{}
}


type NotificationSender interface {
	Send(msg Message)
}


type MockSender struct {}

func (sender *MockSender) Send(msg Message) {
	fmt.Println("MOCK NOTIFICATION")
}


func DispatchNotification(model *models.Model) {
	item := *model
	sender := MockSender{}
	data := item.GetData()
	sender.Send(Message{"None", data})
}