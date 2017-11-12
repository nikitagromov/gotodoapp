package services

import (
	"fmt"
	"todoapp/app/models"
	"time"
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
	fmt.Println("blocked")
	time.Sleep(10 * time.Second)
	sender.Send(Message{"None", data})
}