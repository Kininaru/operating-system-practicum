package main

import (
	"fmt"
	"math"

	"github.com/Kininaru/operating-system-practicum/utils"
)

func opt(memory, queue []int) {
	if index := utils.GetIndex(memory, -1); index != -1 {
		memory[index] = queue[0]
		return
	}

	page2next := make(map[int]int)
	for i, d := range queue[1:] {
		page2next[d] = int(math.Min(float64(page2next[d]), float64(i)))
	}

	latestPage, latestIndex := -1, -1
	for p, i := range page2next {
		if latestIndex < i && utils.GetIndex(memory, p) != -1 {
			latestIndex = i
			latestPage = p
		}
	}

	index := len(memory) - 1
	if latestPage != -1 {
		index = utils.GetIndex(memory, latestPage)
	}

	memory[index] = queue[0]
}

func Opt(page int, queue []int) {
	var count float64
	memory := utils.GetMemory(page)
	for i, p := range queue {
		utils.PrintMemory(memory)
		if utils.GetIndex(memory, p) != -1 {
			continue
		}
		count++
		opt(memory, queue[i:])
	}
	fmt.Println("Final: ")
	utils.PrintMemory(memory)
	fmt.Printf("缺页率: %.3f%%", count/float64(len(queue))*100)
}
