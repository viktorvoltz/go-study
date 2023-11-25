package main

import "fmt"

type messageTosend struct {
	message string
	sender user
	recipient user
}

type user struct {
	name string
	number int
}

func canSendMessage(mToSend messageTosend) bool {
	if mToSend.message == "" {
		return false
	}
	if mToSend.sender.name == "" {
		return false
	}
	if mToSend.recipient.number == 0 {
		return false
	}
	return true
}