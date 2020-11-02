package letcode

//有随机指针的链表复制

type RandListNode struct {
	Val  int
	Next *RandListNode
	Rand *RandListNode
}

func CopyList(node *RandListNode) *RandListNode {
	//先遍历一遍链表变成这样 1--1'--2--2'--3--3'
	current := node //搞一个变量是想复用node
	for current != nil {
		oldNext := current.Next
		newNext := &RandListNode{
			Val:  current.Val,
			Next: oldNext,
			Rand: nil,
		}
		current.Next = newNext
		current = oldNext //直接指向下一个
	}

	//现在开始拆解然后返回复制
	current = node
	for true {
		current.Next.Rand = current.Rand.Next
		current = current.Next.Next
		if current == nil {
			break
		}
	}

	//然后开始恢复
	current = node
	newHead := node.Next
	for true {
		oldNext := current.Next
		newNext := oldNext.Next
		current.Next = newNext
		if newNext != nil {
			oldNext.Next = newNext.Next
		}

		if newNext == nil {
			break
		}

		current = current.Next
	}
	return newHead
}
