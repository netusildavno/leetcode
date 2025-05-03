package main

import "fmt"

// Given an array of integers nums and an integer target, return indices of the two numbers such that they add up to target.
// You may assume that each input would have exactly one solution, and you may not use the same element twice.
// You can return the answer in any order.

// Input: nums = [2,7,11,15], target = 9
// Output: [0,1]
// Explanation: Because nums[0] + nums[1] == 9, we return [0, 1].
func twoSum(nums []int, target int) []int {
	// value -> index map
	m := make(map[int]int)
	for i, val := range nums {
		c := target - val
		if j, ok := m[c]; ok {
			return []int{j, i}
		}
		m[val] = i
	}
	return nil
}

func main() {
	fmt.Println(twoSum([]int{2, 7, 11, 15}, 26))
}
