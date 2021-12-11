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
	//  Leetcode 03
	fmt.Println(lengthOfLongestSubstring("tmmzuxt"))
	//  Leetcode 283
	nums02 := []int{0, 1, 0, 3, 12}
	moveZeroes(nums02)
	fmt.Println(nums02)
	println("------------------------------")
	//  Leetcode

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

//  Leetcode 11. Medium, #1
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

//  Leetcode 03, Medium, #1
func lengthOfLongestSubstring(s string) int {
	i, usedchar, max_count := -1, make(map[string]int), 0
	for j := 0; j < len(s); j++ {
		if _, ok := usedchar[string(s[j])]; ok && usedchar[string(s[j])] > i {
			i = usedchar[string(s[j])]
		} else {
			max_count = max(max_count, j-i)
		}
		usedchar[string(s[j])] = j
	}
	return max_count
}

//  Leetcode 283, Easy ,#1
func moveZeroes(nums []int) {
	i := -1
	for j := 0; j < len(nums); j++ {
		if nums[j] != 0 {
			i++
			nums[i] = nums[j]
		}
		if i != j {
			nums[j] = 0
		}
	}
	return
}
