package main

import "fmt"

func main() {
	L1 := &ListNode{
		2,
		&ListNode{
			4,
			&ListNode{
				3,
				nil,
			},
		},
	}

	L2 := &ListNode{
		5,
		&ListNode{
			6,
			&ListNode{
				4,
				nil,
			},
		},
	}

	fmt.Println(addTwoNumbers(L1, L2))

}

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	l3 := &ListNode{}
	for {
		if l1 == nil {
			l1 = &ListNode{}
		}
		if l2 == nil {
			l2 = &ListNode{}
		}

		f3 := add(l1, l2)

		l1 = l1.Next
		l2 = l2.Next
		if f3.Val == 0 {
			break
		} else if f3.Val >= 10 {
			f4 := &ListNode{
				1,
				nil,
			}
			f3.Val -= 10
			f3.Next = f4
			l3.Next = f3
			break
		}
		l3.Next = f3
	}
	l3 = l3.Next
	return l3
}

// 3种情况 l3.Val == 0  [1,9] [10,20]
func add(l1 *ListNode, l2 *ListNode) *ListNode {
	//处理单节点
	if l1 == nil {
		l1 = &ListNode{}
	}
	if l2 == nil {
		l2 = &ListNode{}
	}
	l3 := &ListNode{}
	l3.Val = l1.Val + l2.Val
	if l3.Val >= 10 {
		if l1.Next != nil {
			l3.Val = l3.Val - 10
			l1.Next.Val++
		} else if l2.Next != nil {
			l3.Val = l3.Val - 10
			l2.Next.Val++
		}
	}

	return l3
}
