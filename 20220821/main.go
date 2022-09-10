package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func getList(lists []*ListNode) *ListNode {
	// write code here
	var res *ListNode
	var temp *ListNode
	b := false
	for i := len(lists) - 1; i >= 0; i-- {
		if !b {
			if lists[i] != nil {
				res = lists[i]
				temp = res
				for temp.Next != nil {
					temp = temp.Next
				}
				b = true
			}
		} else {
			if lists[i] != nil {
				temp.Next = lists[i]
				for temp.Next != nil {
					temp = temp.Next
				}
			} else {
				continue
			}
		}

	}

	return res
}

func main() {
	l1 := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val:  3,
				Next: nil,
			},
		},
	}
	l2 := &ListNode{
		Val: 3,
		Next: &ListNode{
			Val: 4,
		},
	}
	l3 := &ListNode{
		Val: 6,
		Next: &ListNode{
			Val: 7,
			Next: &ListNode{
				Val:  8,
				Next: nil,
			},
		},
	}
	list := getList([]*ListNode{l1, l2, l3})
	fmt.Println(list)
}
