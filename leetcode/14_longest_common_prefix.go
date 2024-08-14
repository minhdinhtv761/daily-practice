package leetcode

import "strings"

// longestCommonPrefix searches prefix by horizontal scanning
// problem: https://leetcode.com/problems/longest-common-prefix/description/
/* TODO:
- searches prefix by vertical scanning
- searches prefix by divide and conquer search
- searches prefix by binary search
*/
func longestCommonPrefix(strs []string) string {
	var res string

	if len(strs) > 0 {
		res = strs[0]
		for i := 1; i < len(strs); i++ {
			for len(res) > 0 {
				if strings.HasPrefix(strs[i], res) {
					break
				}
				res = res[:len(res)-1]
			}
			if len(res) == 0 {
				break
			}
		}
	}

	return res
}
