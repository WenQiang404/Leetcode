/*
*@Author: Wenqiang
*@Date:   2023/5/10
*@Time:   20:38
 */
package main

type Listnode struct {
	Val  int
	Next *Listnode
}

// 自己的思路
func reverseList(head *Listnode) *Listnode {
	headTwo := head
	array := make([]int, 0)
	for head.Next != nil {
		array = append(array, head.Val)
		head = head.Next
	}
	for i := 0; i < len(array); i++ {
		headTwo.Val = array[i]
		headTwo = headTwo.Next
	}
	return headTwo

}

// 大佬的思路1:双指针
func reverseList2(head *ListNode) *ListNode {
	var pre *ListNode
	cur := head
	for cur != nil {
		next := cur.Next
		cur.Next = pre
		pre = cur
		cur = next

	}
	return pre
}

// 递归法：
func reverseList3(head *ListNode) *ListNode {
	return help(nil, head)
}
func help(pre, head *ListNode) *ListNode {
	if head == nil {
		return pre
	}
	next := head.Next
	head.Next = pre
	return help(head, next)
}

//func main() {
//	node2 := Listnode{Val: 3, Next: nil}
//	node1 := Listnode{Val: 2, Next: &node2}
//	node := reverseList(&node1)
//	fmt.Print(node)
//}
