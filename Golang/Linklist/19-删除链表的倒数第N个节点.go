/*
给你一个链表，删除链表的倒数第 n 个结点，并且返回链表的头结点。
*/
package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

// 获取链表长度
func getLength(head *ListNode) int {
	length := 1
	for head.Next != nil {
		length += 1
		head = head.Next
	}
	return length
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	headCopy := head
	length := getLength(head)
	deleteIndex := length - n + 1 //deleteIndex = 4

	if length == 1 {
		return head
	} else if deleteIndex == 1 {
		//删除第一个节点
		head.Next = head.Next.Next
		return head
	}
	for j := 1; j < deleteIndex-1; j++ {
		head = head.Next
	}
	//此时head指向待删除节点的前一个节点
	head.Next = head.Next.Next
	return headCopy
}

func main() {

	list4 := ListNode{Val: 4, Next: nil}
	list5 := ListNode{Val: 5, Next: &list4}
	list6 := ListNode{Val: 6, Next: &list5}
	list7 := ListNode{Val: 7, Next: &list6}
	fmt.Println(getLength(&list7))
	fmt.Printf("%+v", removeNthFromEnd2(&list7, 1))

}

//**********************answer**********************
/*
采取双重遍历肯定是可以解决问题的，但题目要求我们一次遍历解决问题，那我们的思路得发散一下。
我们可以设想假设设定了双指针 p 和 q 的话，当 q 指向末尾的 NULL，p 与 q 之间相隔的元素个数为 n 时，
那么删除掉 p 的下一个指针就完成了要求。
设置虚拟节点 dummyHead 指向 head
设定双指针 p 和 q，初始都指向**虚拟节点 dummyHead**
移动 q，直到 p 与 q 之间相隔的元素个数为 n
同时移动 p 与 q，直到 q 指向的为 NULL
将 p 的下一个节点指向下下个节点

作者：cxywushixiong
链接：https://leetcode.cn/problems/remove-nth-node-from-end-of-list/solution/dong-hua-tu-jie-leetcode-di-19-hao-wen-ti-shan-chu/
来源：力扣（LeetCode）
著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。
*/
func removeNthFromEnd2(head *ListNode, n int) *ListNode {
	if head.Next == nil {
		return nil
	}
	n = -n
	left, right := head, head
	for {
		right = right.Next
		if right == nil {
			break
		}
		if n >= 0 {
			left = left.Next
		} else {
			n++
		}

	}
	if n < 0 {
		return head.Next
	}
	left.Next = left.Next.Next
	return head
}
