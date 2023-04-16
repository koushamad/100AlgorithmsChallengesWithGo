package MagicalString

//A string is considered to be magical if all of its characters are the same.
//You are given a string str of length N containing only lower-case alphabets a to z . You are allowed to do at max k operations on it. In each operation, you can replace any character of the given string with any other character. After at max k operations, what can be the maximum length of its substring which is also a magical string.
//
//Constraints
//1 <= N <= 5 * 10^5
//str[i] = 'a' to 'z'
//0 <= k <= 5 * 10^5
//
//Example
//str = "abaab"
//k = 1
//Output = 4
//Here we can replace 2nd character with 'a'. Resultant string will be "aaaab" and here maximum length of magical substring is 4 ("aaaa")

// O(N)
func maxMagicalSubstring(s string, k int) int {
	charCount := make([]int, 26)
	maxLength := 0
	n := len(s)

	for left, right := 0, 0; right < n; right++ {
		charIndex := int(s[right] - 'a')
		charCount[charIndex]++

		maxCharCount := 0
		for i := 0; i < 26; i++ {
			maxCharCount = max(maxCharCount, charCount[i])
		}

		if right-left+1-maxCharCount > k {
			charCount[int(s[left]-'a')]--
			left++
		} else {
			maxLength = max(maxLength, right-left+1)
		}
	}

	return maxLength
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
