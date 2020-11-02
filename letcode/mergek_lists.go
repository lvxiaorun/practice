package letcode

import (
	"github.com/lvxiaorun/logger"
	"go.uber.org/zap"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func MergeKLists(lists []*ListNode) *ListNode {
	l := len(lists)
	if l%2 != 0 {
		lists = append(lists, nil)
	}
	var result *ListNode
	var args []*ListNode
	for i := 0; i < l; {
		result = MergeTwo(lists[i], lists[i+1])

		logger.Debug("mw", zap.Any("mto", result))
		args = append(args, result)
		i = i + 2
	}
	if len(args) >= 2 {
		return MergeKLists(args)
	}
	return result

}

func MergeTwo(list1 *ListNode, list2 *ListNode) *ListNode {

	var result = &ListNode{
		Val:  0,
		Next: nil,
	}

	item := result
	for list1 != nil && list2 != nil {
		if list1.Val > list2.Val {
			item.Next = list2
			list2 = list2.Next
		} else {
			item.Next = list1
			list1 = list1.Next
		}
		item = item.Next
	}
	if list1 == nil {
		item.Next = list2
		return result.Next
	}
	if list2 == nil {
		item.Next = list1
		return result.Next
	}
	return result.Next
}
