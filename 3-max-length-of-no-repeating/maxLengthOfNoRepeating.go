package main

import "fmt"

func main() {
	fmt.Println(maxLengthOfNoRepeating("abcabcbb"))
	fmt.Println(maxLengthOfNoRepeating("aabbccdef"))
	fmt.Println(maxLengthOfNoRepeating("灰化肥挥发发灰会发黑"))
}

func maxLengthOfNoRepeating(s string) int {
	maxLength := 0
	lastOccur := make(map[rune]int)
	start := 0
	for i, c := range []rune(s) {
		if lastIdx, ok := lastOccur[c]; ok && lastIdx + 1 > start {
			start = lastIdx + 1
		}
		length := i - start + 1
		if length > maxLength {
			maxLength = length
		}
		lastOccur[c] = i
	}
	return maxLength
}
