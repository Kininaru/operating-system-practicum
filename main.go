package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

func main() {
	fmt.Println(os.Args)
	page := 4
	strategy := "fifo"
	var err error

	startIndex := 1
	for i, arg := range os.Args {
		switch arg {
		case "-s":
			strategy = os.Args[i+1]
			startIndex = i + 2
		case "-p":
			page, err = strconv.Atoi(os.Args[i+1])
			if err != nil {
				panic(err)
			}
			startIndex = i + 2
		}
	}

	queue := make([]int, len(os.Args)-startIndex)
	index := 0
	for _, arg := range os.Args[startIndex:] {
		num, err := strconv.Atoi(arg)
		if err != nil {
			panic(err)
		}
		queue[index] = num
		index++
	}

	switch strategy {
	case "fifo":
		Fifo(page, queue)
	case "opt":
		Opt(page, queue)
	case "lru":
		Lru(page, queue)
	default:
		panic(errors.New(fmt.Sprintf("Unknown strategy: %s", strategy)))
	}
}
