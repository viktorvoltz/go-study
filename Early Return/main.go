package main

import (
	"errors"
	"fmt"
)

func divide(dividend, divisor int) (int, error) {
	if divisor == 0 {
		fmt.Println("attempted div by zero")
		return 0, errors.New("can't divide by zero")
	}
	return dividend / divisor, nil
}
