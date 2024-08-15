package leetcode

func lemonadeChange(bills []int) bool {
	billCount := make(map[int]int)
	for _, v := range bills {
		billCount[v]++
	}
	if billCount[10] > billCount[5] {
		return false
	}
	billCount[5] -= billCount[10]
	if billCount[20] > billCount[5] {
		return false
	}
	return billCount[5] >= billCount[20]+(billCount[20]-billCount[10])*2
}
