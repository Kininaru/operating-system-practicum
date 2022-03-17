package main

import (
	"errors"
	"fmt"
)

func main() {
	queue := []int{1, 2, 3, 0, 3, 5, 6, 4, 9}
	page := 4
	policy := "opt"
	
	switch policy {
	case "fifo":
		FirstInFirstOut(page, queue)
	case "opt":
		Opt(page, queue)
	default:
		panic(errors.New(fmt.Sprintf("Unknown policy: %s", policy)))
	}
}
