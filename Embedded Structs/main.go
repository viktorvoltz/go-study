package main

import "fmt"

type sender struct {
	user
	rateLimit int
}

type user struct {
	name string
	number int
}

func test(s sender) {
	fmt.Println("sender name: " s.name)
	fmt.Println("sender number: " s.number)
}

func main() {
	test(sender {
		rateLimit: 30000,
		user: user{
			name: "deb",
			number: 9838938884,
		}
	}),
	test(sender {
		rateLimit: 50000,
		user: user{
			name: "deby",
			number: 9834568884,
		}
	}),
}