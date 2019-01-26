package t2_add_two_numbers

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

type ListNode struct {
	Val int
	Next *ListNode
}
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	if nil == l1 {
		return l2
	}
	if nil == l2 {
		return l1
	}
	h := &ListNode{}
	p, p1, p2 := h, l1, l2
	// 进位
	j := 0
	for nil != p1 && nil != p2 {
		s := p1.Val + p2.Val + j
		// 取个位
		p.Val = s%10
		// 取进位
		j = s/10

		p1, p2 = p1.Next, p2.Next
		// 是否要继续创建
		if nil != p1 || nil != p2 {
			p.Next = &ListNode{}
			p = p.Next
		}
	}

	if p2 != nil { p1 = p2}

	// 处理长链
	for nil != p1 {
		s := p1.Val + j
		p.Val = s%10
		j = s/10

		p1 = p1.Next
		if nil != p1 {
			p.Next = &ListNode{}
			p = p.Next
		}
	}

	// 处理最后的进位
	if j > 0 {
		p.Next = &ListNode{}
		p = p.Next
		p.Val = j
	}

	return h
}