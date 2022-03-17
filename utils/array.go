package utils

import "fmt"

func GetIndex(array []int, target int) int {
	for index, data := range array {
		if data == target {
			return index
		}
	}
	return -1
}

func PrintMemory(memory []int) {
	fmt.Println(memory)
}

func GetMemory(page int) []int {
	memory := make([]int, page)
	for index := range memory {
		memory[index] = -1
	}
	return memory
}
