package letcode

import (
	"fmt"
	"strings"
)

func Trap(height []int) int {
	var area int //总面积
	for true {
		var count int //每一层1的数量
		var str string
		for i := range height {
			if height[i] > 0 {
				height[i] = height[i] - 1
				str = str + "1"
				count++
			} else {
				str = str + "0"
			}
		}
		if count <= 1 {
			break
		}
		area += areaSignle(str)
	}
	return area
}

//计算单层的算法
func areaSignle(str string) int {
	fmt.Println(str)
	//两个1中间的所有的0的个数
	first := strings.Index(str, "1")
	last := strings.LastIndex(str, "1")

	return strings.Count(str[first:last+1], "0")
}

func Trap2(height []int) int {
	l := len(height)
	if l == 0 {
		return 0
	}
	var area int
	left, right := 0, l-1
	maxLeft, maxRight := height[left], height[right]
	for left != right {
		if maxLeft < maxRight {
			left++
			if height[left] <= maxLeft {
				area += maxLeft - height[left]
			} else {
				maxLeft = height[left]
			}
		} else {
			right--
			if height[right] <= maxRight {
				//边上的大于中间的
				area += maxRight - height[right]
			} else {
				maxRight = height[right]
			}
		}
	}
	return area
}
