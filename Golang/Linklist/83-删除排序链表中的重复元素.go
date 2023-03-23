/*
给定一个已排序的链表的头 head ， 删除所有重复的元素，使每个元素只出现一次 。返回 已排序的链表 。
*/
package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func deleteDuplicates(head *ListNode) *ListNode {

	p := head

	if p == nil {
		return nil
	}

	for p.Next != nil {
		if p.Val == p.Next.Val {
			p.Next = p.Next.Next

		} else {
			p = p.Next
		}
	}
	return head
}

func PrintList(head *ListNode) {
	for head != nil {
		fmt.Println(head.Val)
		head = head.Next
	}
}
func main() {
	list4 := ListNode{Val: 1, Next: nil}
	list5 := ListNode{Val: 1, Next: &list4}
	list6 := ListNode{Val: 1, Next: &list5}
	list7 := ListNode{Val: 1, Next: &list6}
	deleteDuplicates(&list7)
	PrintList(&list7)
}
