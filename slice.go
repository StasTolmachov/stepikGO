package main

import "fmt"

func main() {
	list := make([]int, 5, 10)
	fmt.Println(len(list), cap(list))
	fmt.Println(list)
}
