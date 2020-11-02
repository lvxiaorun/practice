package letcode

import (
	"github.com/lvxiaorun/logger"
	"go.uber.org/zap"
	"testing"
)

func TestMergeTwo(t *testing.T) {
	l1 := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val:  4,
			Next: nil,
		},
	}

	l2 := &ListNode{
		Val: 2,
		Next: &ListNode{
			Val:  8,
			Next: nil,
		},
	}
	logger.Debug("tt", zap.Any("re", MergeTwo(l1, l2)))
}

func TestMergeKLists(t *testing.T) {
	l1 := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val:  4,
			Next: nil,
		},
	}

	l2 := &ListNode{
		Val: 2,
		Next: &ListNode{
			Val:  3,
			Next: nil,
		},
	}
	l3 := &ListNode{
		Val: 2,
		Next: &ListNode{
			Val:  5,
			Next: nil,
		},
	}

	l4 := &ListNode{
		Val: 3,
		Next: &ListNode{
			Val:  8,
			Next: nil,
		},
	}
	l := []*ListNode{l1, l2, l3, l4}
	logger.Debug("tt", zap.Any("re", MergeKLists(l)))
}
