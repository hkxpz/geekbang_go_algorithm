package Week01

//ListNode 链表
type ListNode struct {
	Val  int
	Next *ListNode
}

//MergeTwoLists 21. 合并两个有序链表
func MergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	prevHead := &ListNode{0, nil}
	prev := prevHead
	for l1 != nil || l2 != nil {
		if l2 == nil || (l1 != nil && l1.Val <= l2.Val) {
			prev.Next = l1
			l1 = l1.Next
		} else {
			prev.Next = l2
			l2 = l2.Next
		}
		prev = prev.Next
	}
	return prevHead.Next
}

//PlusOne 66. 加一
func PlusOne(digits []int) []int {
	for i := len(digits) - 1; i >= 0; i-- {
		if digits[i] != 9 {
			digits[i]++
			return digits
		} else {
			digits[i] = 0
		}
	}

	digits[0] = 1
	digits = append(digits, 0)
	return digits
}
