package leetcode

// lemonadeChange uses greedy to check all the bills
// problem: https://leetcode.com/problems/lemonade-change/description/
func lemonadeChange(bills []int) bool {
	counts := make(map[int]int)
	for _, v := range bills {
		counts[v]++
		switch v {
		case 10:
			if counts[5] != 0 {
				counts[5]--
				continue
			}
			return false

		case 20:
			if counts[5] != 0 && counts[10] != 0 {
				counts[10]--
				counts[5]--
				continue
			}
			if counts[5] > 2 {
				counts[5] -= 3
				continue
			}
			return false

		default:
			// do nothing
		}
	}
	return true
}
