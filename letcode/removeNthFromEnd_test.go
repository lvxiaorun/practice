package letcode

import (
	"testing"
)

func TestRemoveNthFromEnd(t *testing.T) {
	l := &ListNode{1, &ListNode{
		Val: 2,
		Next: &ListNode{
			Val: 3,
			Next: &ListNode{
				Val:  4,
				Next: nil,
			},
		},
	}}
	//logger.Debug("ttt",zap.Any("res",RemoveNthFromEnd(l,3)))
	re := RemoveNthFromEnd2(l, 2)

	t.Log(re.Next)
	t.Log(re.Next.Next)
	t.Log(re.Next.Next.Next)
}
