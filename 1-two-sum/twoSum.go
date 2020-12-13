package main

import "fmt"

func main() {
	nums := []int{11, 7, 2, 5}
	result := towSum(nums, 9)
	fmt.Println(result)
}

func towSum(nums []int, target int) []int {
	var diffrence int
	numsMap := make(map[int]int)
	for i, n := range nums {
		numsMap[n] = i
		diffrence = target - n
		if idx, ok := numsMap[diffrence]; ok {
			return []int{idx, i}
		}
	}
	return []int{}
}
