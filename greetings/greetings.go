package greetings

import (
	"errors"
	"fmt"
)

func Hello(name string) string {
	message := fmt.Sprintf("Hi ,%v.Welcome !", name)
	return message
}
func Hello2(name string) (string, error) {
	if name == "" {
		return "", errors.New("name is empty")
	}

	return fmt.Sprint("Hello,%s!", name), nil
}
