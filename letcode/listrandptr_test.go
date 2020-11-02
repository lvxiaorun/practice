package letcode

import (
	"fmt"
	"testing"
)

func TestCopyList(t *testing.T) {
	l1 := &RandListNode{
		Val:  1,
		Next: nil,
		Rand: nil,
	}
	l2 := &RandListNode{
		Val:  2,
		Next: nil,
		Rand: nil,
	}
	l3 := &RandListNode{
		Val:  3,
		Next: nil,
		Rand: nil,
	}
	l4 := &RandListNode{
		Val:  4,
		Next: nil,
		Rand: nil,
	}
	l1.Next = l2
	l2.Next = l3
	l3.Next = l4

	l1.Rand = l4
	l2.Rand = l3
	l3.Rand = l1
	l4.Rand = l2
	result := CopyList(l1)
	printList(result)
}

func printList(l *RandListNode) {
	for true {
		fmt.Println(l.Val, l.Rand.Val)
		l = l.Next
		if l == nil {
			return
		}
	}
}
