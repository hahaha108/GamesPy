package leetcode

//给出两个 非空 的链表用来表示两个非负的整数。其中，它们各自的位数是按照 逆序 的方式存储的，并且它们的每个节点只能存储 一位 数字。
//
//如果，我们将这两个数相加起来，则会返回一个新的链表来表示它们的和。
//
//您可以假设除了数字 0 之外，这两个数都不会以 0 开头。
//
//来源：力扣（LeetCode）
//链接：https://leetcode-cn.com/problems/add-two-numbers
//著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

//Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

//运行时间 4 ms
func AddTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	nextAdd := 0
	l3 := &ListNode{0,nil}
	firstNode :=l3
	l3.Val = (l1.Val + l2.Val + nextAdd) % 10
	nextAdd = (l1.Val + l2.Val + nextAdd) / 10
	for {
		if l1.Next ==nil && l2.Next ==nil && nextAdd ==0 {
			break
		}
		sumNumber := 0
		if  l1.Next !=nil {
			l1 = l1.Next
			sumNumber += l1.Val
		}
		if l2.Next !=nil {
			l2 = l2.Next
			sumNumber += l2.Val
		}
		sumNumber += nextAdd
		l3.Next = &ListNode{sumNumber % 10,nil}
		l3 = l3.Next
		nextAdd = sumNumber / 10
	}
	return firstNode
}
