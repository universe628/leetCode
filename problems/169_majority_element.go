package problems

/*
Given an array nums of size n, return the majority element.

The majority element is the element that appears more than ⌊n / 2⌋ times. You may assume that the majority element always exists in the array.



Example 1:

Input: nums = [3,2,3]
Output: 3
Example 2:

Input: nums = [2,2,1,1,1,2,2]
Output: 2


Constraints:

n == nums.length
1 <= n <= 5 * 104
-109 <= nums[i] <= 109

*/

// using Boyer–Moore majority vote algorithm
func majorityElement(nums []int) int {
	var count, target int
	for i := 0; i < len(nums); i++ {
		if count == 0 {
			target = nums[i]
			count = 1
		} else if nums[i] == target {
			count++
		} else {
			count--
		}
	}
	return target
}
