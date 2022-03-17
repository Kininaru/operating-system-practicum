package main

import (
	"errors"
	"fmt"
)

func main() {

	queue := []int{1, 2, 3, 0, 3}
	page := 4
	policy := "fifo"

	for _, data := range queue {
		if data >= page || data < 0 {
			panic(errors.New(fmt.Sprintf("Page index out of bounds. The error index is %d. ", data)))
		}
	}

	switch policy {
	case "fifo":
		FirstInFirstOut(page, queue)
	default:
		panic(errors.New(fmt.Sprintf("Unknown policy: %s", policy)))
	}
}
