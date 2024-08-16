package easy

// romanToInt uses hashmap to calculate the convert roman to int
// problem: https://leetcode.com/problems/roman-to-integer/description/
func romanToInt(s string) int {
	symbol2ValueMap := map[int32]int{
		'I': 1,
		'V': 5,
		'X': 10,
		'L': 50,
		'C': 100,
		'D': 500,
		'M': 1000,
	}
	result := 0
	tmp := 0
	for _, c := range s {
		v := symbol2ValueMap[c]
		result += v
		if tmp < v {
			result -= 2 * tmp
		}
		tmp = v
	}
	return result
}
