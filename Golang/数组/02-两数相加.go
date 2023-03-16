package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

// 创建保存数字的栈，保存个十百千位
type stack struct {
	Top  int
	List []int
}

// 定义压栈方法
func (s *stack) push(x int) {
	s.Top++
	s.List = append(s.List, x)
}

// 定义出栈方法
func (s *stack) pop() (int, bool) {
	if s.Top == -1 {
		return 0, false
	}
	v := s.List[s.Top]
	s.Top--
	return v, true
}

var i = 0

// 目前有问题
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	//l1Length := l1.getLength()		//保存链表1的长度
	//l2Length := l2.getLength()		//保存链表2的长度
	stack1 := &stack{Top: -1}
	stack2 := &stack{Top: -1}
	stack3 := &stack{Top: -1} //保存结果
	flag := 0                 //是否进位

	for l1.Next != nil {
		stack1.push(l1.Val)
	}
	for l2.Next != nil {
		stack2.push(l2.Val)
	}

	//长度相等，否则不等
	if stack1.Top == stack2.Top {
		for stack2.Top != -1 {

			result := stack1.List[stack1.Top] + stack2.List[stack2.Top]
			if result < 10 {
				stack3.push(result + flag)
			} else {
				stack3.push(result % 10)
				flag = 1
			}
			stack1.Top--
			stack2.Top--
		}
	} else {
		shortStack := 0
		if stack1.Top > stack2.Top {
			shortStack = stack2.Top
			i = 1
		} else {
			shortStack = stack1.Top
			i = 2
		}
		for shortStack != -1 {
			result := stack1.List[stack1.Top] + stack2.List[stack2.Top]
			if result < 10 {
				stack3.push(result + flag)
			} else {
				stack3.push(result % 10)
				flag = 1
			}
			stack1.Top--
			stack2.Top--
		}
		for stack2.Top != -1 {
			result2 := stack2.List[stack2.Top] + flag
			if result2 > 10 {
				stack3.push(result2 % 10)
				stack3.push(result2 / 10)
			} else {
				stack3.push(result2)
			}
			if i == 1 {
				stack3.push(stack1.List[stack1.Top])
			} else {
				stack3.push(stack2.List[stack2.Top])
			}
		}
	}

	//输出到链表
	finalList := &ListNode{}
	for stack3.Top != -1 {
		finalList.Val = stack3.List[stack3.Top]
		finalList.Next = finalList.Next.Next
	}
	return finalList
}

// 官方解答
func addTwoNumber(l1 *ListNode, l2 *ListNode) (head *ListNode) {
	//定义一个尾结点，或者可以理解为临时节点
	var tail *ListNode
	//余数carry
	carry := 0
	//依次遍历两个链表，只要元素不为空就进行下一步
	for l1 != nil || l2 != nil {
		//定义两个变量存储各个节点的值
		n1, n2 := 0, 0
		//从第一个链表开始
		if l1 != nil {
			//把每个节点的值赋给n1
			n1 = l1.Val
			//节点后移
			l1 = l1.Next
		}
		//l2同上
		if l2 != nil {
			n2 = l2.Val
			l2 = l2.Next
		}
		//此时是两个链表第一个元素的和 + 余数
		sum := n1 + n2 + carry
		//sum%10是节点的当前值，如果是10,取余后当前节点值为0，sum/10是求十位的那个数
		sum, carry = sum%10, sum/10
		//此时申请一个新的链表存储两个链表的和
		if head == nil {
			//申请新的链表
			head = &ListNode{Val: sum}
			//这一步是为了保持头结点不变的情况下指针可以右移，所以说tail相当于临时节点，理解成尾节点也可以，因
			//为此时新链表中只有一个节点，所以头结点和尾结点都指向同一个元素。
			tail = head
		} else {
			//第二个节点后开始逐渐往尾结点增加元素
			tail.Next = &ListNode{Val: sum}
			tail = tail.Next
		}
	}
	//把最后一位的余数加到链表最后。
	if carry > 0 {
		tail.Next = &ListNode{Val: carry}
	}
	return
}

func main() {
	list3 := ListNode{Val: 3, Next: nil}
	list2 := ListNode{Val: 2, Next: &list3}
	list1 := ListNode{Val: 1, Next: &list2}

	list4 := ListNode{Val: 3, Next: nil}
	list5 := ListNode{Val: 2, Next: &list4}
	list6 := ListNode{Val: 1, Next: &list5}
	list7 := ListNode{Val: 1, Next: &list6}

	fmt.Println(addTwoNumber(&list1, &list7))
}
