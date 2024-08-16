package easy

// twoSum uses hashmap with one pass to solve the problem
// problem: https://leetcode.com/problems/two-sum/description/
func twoSum(nums []int, target int) []int {
	args := make(map[int]int)
	res := make([]int, 2)

	for idx1, vl1 := range nums {
		if idx2, ok := args[target-vl1]; ok {
			res[0] = idx1
			res[1] = idx2
		}
		args[vl1] = idx1
	}
	return res
}
