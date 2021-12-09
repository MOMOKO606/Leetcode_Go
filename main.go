package main

import "fmt"

func main() {
	fmt.Println("Hello, Leetcode!")

	nums01 := []int{2, 7, 11, 15}
	target := 9

	fmt.Println(twoSum(nums01, target))
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
