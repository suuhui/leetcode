package main

import (
	"fmt"
)

func main() {
	fmt.Println(findMedianSortedArray1([]int{1,2,3}, []int{4,5,6}))
	fmt.Println(findMedianSortedArray1([]int{1,3,5}, []int{2,4,6}))
	fmt.Println(findMedianSortedArray1([]int{1,3,3}, []int{2,3,6}))
	fmt.Println(findMedianSortedArray1([]int{1,3,8}, []int{2,4,5,6}))

	fmt.Println("--------------------------------------------------")

	fmt.Println(findMedianSortedArray([]int{1,2,3}, []int{4,5,6}))
	fmt.Println(findMedianSortedArray([]int{1,3,5}, []int{2,4,6}))
	fmt.Println(findMedianSortedArray([]int{1,3,3}, []int{2,3,6}))
	fmt.Println(findMedianSortedArray([]int{1,3,8}, []int{2,4,5,6}))
}

func findMedianSortedArray(arr1, arr2 []int) float64 {
	len1, len2 := len(arr1), len(arr2)
	//将奇数，偶数情况合一，奇数情况下，会获取两次同样的值
	left, right := (len1 + len2 + 1) / 2, (len1 + len2 + 2) / 2
	leftVal := getKth(arr1, 0, len1, arr2, 0, len2, left)
	rightVal := getKth(arr1, 0, len1, arr2, 0, len2, right)
	return float64(leftVal + rightVal) / 2.0
}

func getKth(arr1 []int, start1 int, end1 int, arr2 []int, start2 int, end2 int, k int) int {
	len1 := end1 - start1
	len2 := end2 - start2
	//保证后续len1 < len2 从而长度为0的一定是arr1
	if len1 > len2 {
		return getKth(arr2, start2, end2, arr1, start1, end1, k)
	}
	if len1 == 0 {
		return arr2[start2 + k - 1]
	}
	if k == 1 {
		return intMin(arr1[start1], arr2[start2])
	}
	i := start1 + intMin(len1, k / 2) - 1
	j := start2 + intMin(len2, k / 2) - 1
	if arr1[i] > arr2[j] {
		return getKth(arr1, start1, end1, arr2, j + 1, end2, k - (j - start2 + 1))
	} else {
		return getKth(arr1, i + 1, end1, arr2, start2, end2, k - (i - start1 + 1))
	}
}

func intMin(i, j int) int {
	if i > j {
		return j
	}
	return i
}

func findMedianSortedArray1(arr1, arr2 []int) float64 {
	len1, len2 := len(arr1), len(arr2)
	length := len1 + len2
	endPos := length / 2
	left, right := -1, -1
	start1, start2 := 0, 0
	for i := 0; i <= endPos ; i++ {
		left = right
		if start1 < len1 && (start2 >= len2 || arr1[start1] < arr2[start2]) {
			right = arr1[start1]
			start1++
		} else {
			right = arr2[start2]
			start2++
		}
	}
	if length & 1 == 0 {
		return float64(left + right) / 2.0
	} else {
		return float64(right)
	}
}

