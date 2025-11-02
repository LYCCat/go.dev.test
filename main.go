package main

import "fmt"

func main() {
	num := []int{1, 2, 3, 4, 5}
	num = append(num, num...)
	fmt.Println(num)
}
