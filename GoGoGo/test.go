package main

import (
	"awesomeProject/leetcode"
	"fmt"
)

//func function(a *int) {
//	*a += 100
//}
func main() {
	//l1_3 := &leetcode.ListNode{3,nil}
	//l1_2 := &leetcode.ListNode{4,l1_3}
	//l1 := &leetcode.ListNode{2,l1_2}
	//
	//l2_3 := &leetcode.ListNode{4,nil}
	//l2_2 := &leetcode.ListNode{6,l2_3}
	//l2 := &leetcode.ListNode{5,l2_2}

	//l1 := &leetcode.ListNode{5,nil}
	//l2 := &leetcode.ListNode{5,nil}
	//
	//r := leetcode.AddTwoNumbers(l1,l2)
	//fmt.Print(r)
	//fmt.Print(r.Next)
	//fmt.Print(r.Next.Next)

	//s := "dvdfxascascas"
	//c := leetcode.LengthOfLongestSubstring(s)

	//c := leetcode.LongestPalindrome("babadada")
	//fmt.Print(c)

	ll := [][]int{[]int{1,2,3},[]int{4,5,6},[]int{7,8,9}}
	leetcode.Rotate(ll)
	fmt.Print(ll)
}