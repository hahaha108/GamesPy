package leetcode

//给定一个字符串，请你找出其中不含有重复字符的 最长子串 的长度。

//双重循环运行效率低,淘汰
//func LengthOfLongestSubstring(s string) int {
//	if len(s) == 0{
//		return 0
//	}
//	type void struct{}
//	maxLen := 1
//	for i := 0; i < len(s); i++ {
//		subStrLen := 1
//		subStrMap := make(map[string]void)
//		subStrMap[string(s[i])] = void{}
//		for j := i+1; j < len(s); j++ {
//			if _, ok := subStrMap[string(s[j])]; ok {
//				subStrLen = 1
//				continue
//			} else {
//				subStrLen += 1
//				subStrMap[string(s[j])] = void{}
//				if subStrLen > maxLen {
//					maxLen = subStrLen
//				}
//			}
//		}
//	}
//	return maxLen
//}

//运行时间 412 ms 依旧比较慢
func LengthOfLongestSubstring(s string) int {
	if len(s) == 0 {
		return 0
	}
	maxLen := 0
	subLen := 0
	subStrMap := make(map[string]int)
	for i := 0; i < len(s); i++ {
		if _, ok := subStrMap[string(s[i])]; ok {
			subLen = 0
			i = subStrMap[string(s[i])]
			subStrMap = make(map[string]int)
			continue
		} else {
			subLen += 1
			subStrMap[string(s[i])] = i
			if subLen > maxLen {
				maxLen = subLen
			}
		}
	}
	return maxLen
}
