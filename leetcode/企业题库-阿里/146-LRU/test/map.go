package main

import "fmt"

func main() {
	t := make(map[int]int, 1)
	t[1] = 2
	t[1] = 3
	fmt.Println(t[1])
}