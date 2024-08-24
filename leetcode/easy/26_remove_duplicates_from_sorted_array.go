package easy

// removeDuplicates uses two pointers to iterate through the input and update its elements
// problem: https://leetcode.com/problems/remove-duplicates-from-sorted-array/description/
func removeDuplicates(nums []int) int {
	k := 1
	for idx := 1; idx < len(nums); idx++ {
		if nums[idx-1] != nums[idx] {
			nums[k] = nums[idx]
			k++
		}
	}
	return k
}
