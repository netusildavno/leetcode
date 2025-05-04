package main

import "fmt"

// Given a string s, find the length of the longest substring without duplicate characters.
// Input: s = "abcabcbb"
// Output: 3
// Explanation: The answer is "abc", with the length of 3.

// Naive solution, O(N^2) in worst case.
// Better solution TBD soon.
func lengthOfLongestSubstring(s string) int {
	max := 0
	for l := 0; l < len(s); l++ {
		hits := make(map[string]struct{})
		count := 0
		for r := l; r < len(s); r++ {
			v := string(s[r])
			if _, ok := hits[v]; ok {
				break
			}
			hits[v] = struct{}{}
			count++
			if count > max {
				max = count
			}
		}
	}
	return max
}

func main() {
	fmt.Println(lengthOfLongestSubstring("aaaabbqweraaa"))
}
