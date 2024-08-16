package medium

import "math"

// maxDistance use greedy to get the max distance at each sub array
// problem: https://leetcode.com/problems/maximum-distance-in-arrays/description/
func maxDistance(arrays [][]int) int {
	m := len(arrays)
	minNum, maxNum := arrays[0][0], arrays[0][len(arrays[0])-1]
	result := math.MinInt
	for i := 1; i < m; i++ {
		subArray := arrays[i]
		l := len(subArray)
		result = max(result, subArray[l-1]-minNum, maxNum-subArray[0])
		minNum = min(minNum, subArray[0])
		maxNum = max(maxNum, subArray[l-1])
	}
	return result
}
