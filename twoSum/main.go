package main

import "fmt"

func sumTwo(nums []int, target int) []int {
	for i := 0; i < len(nums); i++ {
		for j := 1; j < len(nums); j++ {
			if (nums[i] + nums[j]) == target {
				return []int{i, j}
			}
		}
	}
	return []int{}
}

func main() {
	fmt.Println("Hello World")
}
