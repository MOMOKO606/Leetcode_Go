package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hello, Leetcode!")

	nums01 := []int{2, 7, 11, 15}
	target := 9
	//  Leetcode 01
	fmt.Println(twoSum(nums01, target))
	//  Leetcode 11
	fmt.Println(maxArea([]int{1, 8, 6, 2, 5, 4, 8, 3, 7}))
}

//  Leetcode 1. Two Sum(Easy), #1
func twoSum(nums []int, target int) (ans []int) {
	hashtable := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		key := target - nums[i]
		if _, ok := hashtable[key]; ok {
			ans = append(ans, hashtable[key])
			ans = append(ans, i)
			break
		} else {
			hashtable[nums[i]] = i
		}
	}
	return
}

func maxArea(height []int) int {

	i, j, area := 0, len(height)-1, 0

	for i < j {
		area = max(area, (j-i)*min(height[i], height[j]))
		if height[i] < height[j] {
			i++
		} else {
			j--
		}
	}
	return area
}

func min(a int, b int) int {
	if a < b {
		return a
	} else {
		return b
	}

}

func max(a int, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}
