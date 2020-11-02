package letcode

import (
	"fmt"
	"math"
	"strings"
)

func MyAtoi(str string) int {
	//去除空格
	str = strings.TrimSpace(str)
	if len(str) == 0 {
		return 0
	}
	mul := 1
	result := 0
	for i, v := range str {
		if v >= '0' && v <= '9' {
			//说明是数字
			result = result*10 + int(v-'0')
			fmt.Println(result)
		} else if v == '-' && i == 0 {
			mul = -1
		} else if v == '+' && i == 0 {
			continue
		} else {
			break
		}
		if result > math.MaxInt32 {
			if mul == -1 {
				return math.MinInt32
			}
			return math.MaxInt32
		}
	}
	return mul * result
}
