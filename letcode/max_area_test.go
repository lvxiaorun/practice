package letcode

import "testing"

func MaxArea(height []int) int {
	l := len(height)
	if l <= 1 {
		return 0
	}
	left := 0
	right := l - 1
	max := 0
	area := 0
	for left != right {
		if height[left] < height[right] {
			area = height[left] * (right - left)
			left++
		} else {
			area = height[right] * (right - left)
			right--
		}
		if area > max {
			max = area
		}
	}
	return max
}

func TestMaxArea(t *testing.T) {
	t.Log(MaxArea([]int{1, 8, 6, 2, 5, 4, 8, 3, 7}))
}
