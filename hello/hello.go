package main

import (
	"example.com/greetings"
	"fmt"
)
import "rsc.io/quote"

func main() {
	fmt.Println("hello world")
	fmt.Println(quote.Go())
	message := greetings.Hello("Gladys")
	fmt.Println(message)

}
