/*
将两个升序链表合并为一个新的 升序 链表并返回。新链表是通过拼接给定的两个链表的所有节点组成的。
*/
package main

//	func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
//		if list1 == nil && list2 == nil {
//			return nil
//		}
//		if list1 == nil {
//			return list2
//		}
//		if list2 == nil {
//			return list1
//		}
//		p, q := list1, list2
//		var head *ListNode
//		var headList *ListNode
//		head.Next = list1
//		for p != nil || q != nil {
//			if q.Val <= p.Val {
//				head.Next = q
//				headList = head
//				q.Next = p
//				q = q.Next
//			} else {
//				p = p.Next
//			}
//		}
//
//
//		return list1
//	}
//
// 合理设置虚拟空节点以及指针复制
func insertNode(list1 *ListNode, list2 *ListNode) *ListNode {
	dummyHead := &ListNode{}
	p := dummyHead
	for list1 != nil && list2 != nil {
		if list1.Val >= list2.Val {
			p.Next = list2
			list2 = list2.Next
		} else {
			p.Next = list1
			list1 = list1.Next
		}
		p = p.Next

	}
	if list1 != nil {
		p.Next = list1
	}
	if list2 != nil {
		p.Next = list2
	}
	return dummyHead.Next
}
