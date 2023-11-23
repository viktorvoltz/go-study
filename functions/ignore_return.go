package main

import "fmt"

func maint() {
	firstname, _ := getNames()
	fmt.Println("welcome to chanel ", firstname)
}

func getNames() (string, string) {
	return "john", "joseph"
}
