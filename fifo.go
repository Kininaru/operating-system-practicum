package main

import (
	"fmt"

	"github.com/Kininaru/operating-system-practicum/utils"
)

func fifo(memory []int, data int) {
	for i := 0; i < len(memory)-1; i++ {
		memory[i] = memory[i+1]
	}
	memory[len(memory)-1] = data
}

func Fifo(page int, queue []int) {
	var count float64
	memory := utils.GetMemory(page)
	for _, d := range queue {
		utils.PrintMemory(memory)
		index := utils.GetIndex(memory, d)
		if index == -1 {
			count++
			fifo(memory, d)
		}
	}
	fmt.Println("Final: ")
	utils.PrintMemory(memory)
	fmt.Printf("缺页率: %5f", count/float64(len(queue)))
}
