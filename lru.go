package main

import (
	"fmt"
	"math"

	"github.com/Kininaru/operating-system-practicum/utils"
)

func lru(memory, lastUsed []int, page, now int) {
	if index := utils.GetIndex(memory, -1); index != -1 {
		memory[index] = page
		lastUsed[index] = now
		return
	}

	index, lastUsedTime := -1, math.MaxInt
	for i, d := range lastUsed {
		if d < lastUsedTime {
			index = i
			lastUsedTime = d
		}
	}

	memory[index] = page
	lastUsed[index] = now
}

func Lru(page int, queue []int) {
	memory := utils.GetMemory(page)
	lastUsed := make([]int, page)
	var count float64
	for i, d := range queue {
		utils.PrintMemory(memory)
		if utils.GetIndex(memory, d) != -1 {
			continue
		}
		lru(memory, lastUsed, d, i)
		count++
	}
	fmt.Println("Final: ")
	utils.PrintMemory(memory)
	fmt.Printf("缺页率: %5f", count/float64(len(queue)))
}
