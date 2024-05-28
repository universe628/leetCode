package problems

/*
You are given an integer array nums. You are initially positioned at the array's first index, and each element in the array represents your maximum jump length at that position.

Return true if you can reach the last index, or false otherwise.



Example 1:

Input: nums = [2,3,1,1,4]
Output: true
Explanation: Jump 1 step from index 0 to 1, then 3 steps to the last index.
Example 2:

Input: nums = [3,2,1,0,4]
Output: false
Explanation: You will always arrive at index 3 no matter what. Its maximum jump length is 0, which makes it impossible to reach the last index.


Constraints:

1 <= nums.length <= 104
0 <= nums[i] <= 105

*/

// my solution
func _canJump(nums []int) bool {
next:
	for i := 0; i < len(nums)-1; i++ {
		if nums[i] == 0 {
			for j := 0; j < i; j++ {
				if nums[j] > i-j {
					continue next
				}
			}
			return false
		}
	}
	return true
}

// a better solution
func canJump(nums []int) bool {
	prev := nums[0]
	for i := 1; i < len(nums); i++ {
		if prev == 0 {
			return false
		}
		//1.21 syntax
		prev = max(prev-1, nums[i])
	}
	return true
}
