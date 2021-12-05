package Week01

//ListNode 链表
type ListNode struct {
	Val  int
	Next *ListNode
}

//MyCircularDeque 双端队列

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

//Constructor 641. 设计循环双端队列

type MyCircularDeque struct {
	Val     int
	Next    *MyCircularDeque
	protect *MyCircularDeque
	cap     int
	size    int
}

func Constructor(k int) MyCircularDeque {
	return MyCircularDeque{protect: &MyCircularDeque{cap: k}}
}

func (this *MyCircularDeque) InsertFront(value int) bool {
	if this.IsFull() {
		return false
	}

	if !this.IsEmpty() {
		this.Next = this
		this.Val = value
		return true
	}

	this.protect.Next = this
	this.protect.size++
	return true
}

func (this *MyCircularDeque) InsertLast(value int) bool {
	if this.IsFull() {
		return false
	}

	if this.IsEmpty() {
		this.protect.Next = this
		this.protect.size++
		return true
	}

	for this.Next != nil {
		this = this.Next
	}

	this.Next = &MyCircularDeque{Val: value}
	this.protect.size++
	return true

}

func (this *MyCircularDeque) DeleteFront() bool {
	if this.IsEmpty() {
		return false
	}

	if this.protect.size == 1 {
		this.protect.Next = nil
		this.protect.size--
		return true
	}

	this.protect.Next = this.protect.Next.Next
	this.protect.size--
	return true
}

func (this *MyCircularDeque) DeleteLast() bool {
	if this.IsEmpty() {
		return false
	}

	for this.Next.Next != nil {
		this = this.Next
	}

	this.Next = nil
	this.protect.size--
	return true
}

func (this *MyCircularDeque) GetFront() int {
	if this.IsEmpty() {
		return -1
	}
	return this.protect.Next.Val
}

func (this *MyCircularDeque) GetRear() int {
	if this.IsEmpty() {
		return -1
	}

	for this.Next != nil {
		this = this.Next
	}

	return this.Val
}

func (this *MyCircularDeque) IsEmpty() bool {
	if this.protect.Next == nil {
		return true
	}
	return false
}

func (this *MyCircularDeque) IsFull() bool {
	if this.protect.cap <= this.protect.size {
		return true
	}
	return false
}
