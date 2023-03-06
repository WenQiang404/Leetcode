package main

type ListNode struct {
	Val  int
	Next *ListNode
}

//方法一：先记录list中所有元素的个数，然后挨个给新数组赋值

func reversePrint_01(head *ListNode) []int {
	count := 0  //record the length of the list
	cur := head //store the list
	for head != nil {
		head = head.Next
		count++

	}
	final_array := make([]int, count)
	for {
		if cur == nil {
			break
		}
		final_array[count-1] = cur.Val
		cur = cur.Next
		count--

	}
	return final_array
}

//方法二 递归解法
func reversePrint_02(head *ListNode) []int {
	if head == nil {
		return []int{}
	}
	return append(reversePrint_02(head.Next), head.Val)
}

//方法三 使用栈
func reversePrint_03(head *ListNode) []int {

	stack := make([]int, 0)
	for head != nil {
		stack = append(stack, head.Val) //遍历入栈
		head = head.Next
	}

	//pop
	var depth_stack = len(stack)
	final_stack := make([]int, depth_stack)
	for i := 0; i < depth_stack; i++ {
		final_stack[i] = stack[depth_stack-1-i]

	}
	return final_stack

}
