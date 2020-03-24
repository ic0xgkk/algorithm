package main

import "fmt"

func exist(board [][]byte, word string) bool {
	t := make(map[string]bool)
	for _, e := range board {
		t[fmt.Sprint(e)] = true
	}
	_, ok := t[word]
	return ok
}