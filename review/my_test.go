package review

import (
	"fmt"
	"testing"
)

func TestMidFind(t *testing.T) {
	a := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	for _, item := range a {
		fmt.Println("run", item)
		t.Log(MidFind(a, item))
	}
}

func TestQuickSort(t *testing.T) {
	a := []int{9, 8, 7, 5, 4, 3, 6}
	fmt.Println(Quick_sort(a))
}

func TestFindClose(t *testing.T) {
	a := []int{0, 1, 2, 3, 5, 6, 7}
	for _, item := range a {
		fmt.Println(FindClose(a, item))
	}
}
