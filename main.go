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
	//  Leetcode 26
	fmt.Println(removeDuplicates([]int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}))
	//  Leetcode 55
	println("------------------------------")
	fmt.Println(canJump([]int{3, 2, 1, 0, 4}))

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

//  Leetcode 283, Easy, #1
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

//  Leetcode 26, Easy, #1
func removeDuplicates(nums []int) int {
	j := -1
	for i := 0; i < len(nums); i++ {
		if j < 0 || nums[i] != nums[i-1] {
			j++
			nums[j] = nums[i]
		}
	}
	return j + 1
}

//  Version that looks longer but faster
//func removeDuplicates(nums []int) int {
//	n := len(nums)
//	if n == 0{
//		return 0
//	}
//	j := 0
//	for i := 1; i < len(nums); i++ {
//		if nums[i] != nums[i - 1]{
//			j++
//			nums[j] = nums[i]
//		}
//	}
//	return j + 1
//}

//  leetcode 55

////  The recursive version with a memo.
//func canJump(nums []int) bool {
//
//	memo := make([]interface{}, len(nums)) //定义长度的slice！
//	memo[0] = true
//
//	return _canJump(nums, memo)
//
//}
//
//func _canJump(nums []int, memo []interface{}) bool {
//	n := len(nums)
//
//	//  等号右侧是check interface memo[n - 1]是否是bool型；
//	//  等号左边ans表示把 memo[n - 1]转化为bool后的值；
//	//  ok表示memo[n - 1]是否为bool型。
//	if ans, ok := memo[n-1].(bool); ok {
//		return ans
//	}
//	for i := 0; i < n-1; i++ {
//		if i + nums[i] >= n-1 {
//			if _canJump(nums[:i+1], memo) {
//				memo[n-1] = true
//				return true
//			}
//		}
//	}
//	memo[n-1] = false
//	return false
//}

////  The iterative version with a memo.
//func canJump( nums []int ) bool{
//	n := len(nums)
//	memo := make([]bool, n)
//	memo[0] = true
//	for i:= 1; i < n; i++{
//		for j:= 0; j < i; j++{
//			if nums[j] + j >= i && memo[j]{
//				memo[i] = true
//				break
//			}
//		}
//		if memo[i] != true{
//			memo[i] = false}
//	}
//	return memo[n - 1]
//}

//  The greedy version.
func canJump(nums []int) bool {
	reach := 0
	n := len(nums)
	for i := 0; i < n; i++ {
		if i > reach {
			return false
		}
		if reach >= n-1 {
			return true
		}
		reach = max(reach, i+nums[i])
	}
	return true
}
