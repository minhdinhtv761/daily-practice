package easy

// isPalindrome reverses the entire number x, then check if new number is equal to the original number
// problem: https://leetcode.com/problems/palindrome-number/description/
// TODO: solve by reversing half of the number x
func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}
	clone := x
	y := 0
	for x != 0 {
		y = y*10 + x%10
		x /= 10
	}
	return clone == y
}
