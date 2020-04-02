package letcode

import (
	"fmt"
	"sort"
	"testing"
)

func ThreeSum(nums []int) [][]int {
	sort.Ints(nums)
	l := len(nums)
	var result [][]int

	var mc = map[string]bool{} //记录已经作为中心处理过
	//- -2 -3 -4 -5 0 1 2 3 4 5
	for i := 1; i < l-1; i++ {
		left := 0
		right := l - 1
		for left != i && right != i {
			sum := nums[i] + nums[left] + nums[right]
			if sum == 0 {
				key := fmt.Sprintf("%d%d%d", nums[left], nums[i], nums[right])
				if mc[key] == false {
					result = append(result, []int{nums[left], nums[i], nums[right]})
					mc[key] = true
				}
				left = left + 1
				right = right - 1
			} else if sum < 0 {
				if nums[i] < 0 && nums[left] < 0 && nums[right] < 0 {
					break
				}
				left = left + 1
			} else {
				if nums[i] > 0 && nums[left] > 0 && nums[right] > 0 {
					break
				}
				right = right - 1
			}
		}
	}
	return result
}

func TestThreeSum(t *testing.T) {
	a := []int{-4, -2, 1, -5, -4, -4, 4, -2, 0, 4, 0, -2, 3, 1, -5, 0}
	fmt.Println(ThreeSum(a))
}
