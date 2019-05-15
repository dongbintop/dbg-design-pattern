package bridge

import "fmt"

type MessageInterface interface {
	SendMessage(message, user string)
}

type MessagePhone struct {
}

func (mp *MessagePhone) SendMessage(message, user string) {
	fmt.Println("phone send message,", message, ",", user)
}

type MessageEmail struct {
}

func (me *MessageEmail) SendMessage(message, user string) {
	fmt.Println("email send message,", message, ",", user)
}

type AbstractMessage struct {
	Message MessageInterface
}

func (am *AbstractMessage) SendMessage(message, user string) {
	am.Message.SendMessage(message, user)
}

type CommonMessage struct {
	*AbstractMessage
}

func (cm *CommonMessage) SendMessage(message, user string) {
	message = "common," + message
	cm.AbstractMessage.SendMessage(message, user)
}

type UrgentMessage struct {
	*AbstractMessage
}

func (um *UrgentMessage) SendMessage(message, user string) {
	message = "urgent," + message
	um.AbstractMessage.SendMessage(message, user)
}

func Test() {
	email := MessageEmail{}
	commonMessage := CommonMessage{&AbstractMessage{&email}}
	commonMessage.SendMessage("hello world", "db")
}
