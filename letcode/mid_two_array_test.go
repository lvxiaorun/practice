package letcode

import (
	"fmt"
	"testing"
)

func TestFindMedianSortedArrays(t *testing.T) {
	a1 := []int{1, 2, 3, 4, 5}
	a2 := []int{6, 7, 8, 9, 10}
	fmt.Println(FindMedianSortedArrays(a1, a2))
}
