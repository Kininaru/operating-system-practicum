package main

import (
	"errors"
	"math"

	"github.com/Kininaru/operating-system-practicum/utils"
)

func lru(memory, lastUsed []int, page, now int) {
	if utils.GetIndex(memory, page) != -1 {
		return
	}

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

	if index == -1 {
		panic(errors.New("invalid index in lru"))
	}

	memory[index] = page
	lastUsed[index] = now
}

func Lru(page int, queue []int) {
	memory := utils.GetMemory(page)
	lastUsed := make([]int, page)

	for i, d := range queue {
		lru(memory, lastUsed, d, i)
		utils.PrintMemory(memory)
	}
}
