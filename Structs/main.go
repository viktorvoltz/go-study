package structs

import "fmt"

type messageToSend struct {
	message     string
	phoneNumber int
}

func test(m messageToSend) {
	fmt.Printf("sending message %s to: %v\n", m.message, m.phoneNumber)
	fmt.Println("=======")
}

func main() {
	test(messageToSend{
		message:     "hello world",
		phoneNumber: 90288938464,
	})
	test(messageToSend{
		message:     "hello nonso",
		phoneNumber: 9028895874,
	})
}
