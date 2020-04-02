package letcode

func RemoveNthFromEnd(head *ListNode, n int) *ListNode {
	//用map实现
	i := 0
	ptr := head
	var mNode = map[int]*ListNode{}
	for ptr != nil {
		mNode[i] = ptr
		ptr = ptr.Next
		i++
	}
	//存入完成 长度是i
	//删除倒数第n个下标是 i-n-1
	//要考虑边界 删除倒数第i个（头）删除倒数第一个（不用特殊处理）
	if n == i {
		return head.Next
	}
	mNode[i-n-1].Next = mNode[i-n].Next
	return head
}

func RemoveNthFromEnd2(head *ListNode, n int) *ListNode {
	//先找到第n个结点
	var nNode = head
	i := 1
	for i < n {
		nNode = nNode.Next
		i++
	}
	//找到n之后 一起向后移动
	if nNode.Next == nil {
		//说明是要找第一个
		return head.Next
	}

	h := head
	for nNode.Next.Next != nil {
		nNode = nNode.Next
		h = h.Next
	}
	h.Next = h.Next.Next
	return head
}
