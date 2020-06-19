package leetcode

//给定一个字符串 s，找到 s 中最长的回文子串。你可以假设 s 的最大长度为 1000。

// 执行用时：116 ms
func LongestPalindrome(s string) string {
	switch len(s) {
	case 0:
		{
			return ""
		}
	case 1:
		{
			return s
		}
	case 2:
		{
			if s[0] == s[1] {
				return s
			} else {
				return string(s[0])
			}
		}
	}
	PalindromeMap := make(map[int]string)
	maxLen := 1
	PalindromeMap[1] = string(s[0])
	if s[0] == s[1] {
		PalindromeMap[2] = s[0:2]
		maxLen = 2
	}
	for i := 2; i < len(s); i++ {
		if s[i] == s[i-1] {
			// 开始检测回文长度
			for j := 0; j < len(s)-i; j++ {
				if i-j < 1 {
					break
				}
				if s[i+j] == s[i-1-j] {
					PalindromeMap[2+j*2] = s[i-1-j : i+j+1]
					if maxLen < 2+j*2 {
						maxLen = 2 + j*2
					}
				} else {
					break
				}
			}
		}
		if s[i] == s[i-2] {
			// 开始检测回文长度
			for j := 0; j < len(s)-i; j++ {
				if i-j < 2 {
					break
				}
				if s[i+j] == s[i-2-j] {
					PalindromeMap[2+j*2+1] = s[i-2-j : i+j+1]
					if maxLen < 2+j*2+1 {
						maxLen = 2 + j*2 + 1
					}
				} else {
					break
				}
			}
		}

	}
	return PalindromeMap[maxLen]
}
