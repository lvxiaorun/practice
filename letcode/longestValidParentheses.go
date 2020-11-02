package letcode

import (
	"container/list"
)

func LongestValidParentheses(s string) int {
	l := list.New()
	l.PushFront(-1)
	var max int
	for i, item := range s {

		if string(item) == "(" {
			l.PushFront(i)
		} else {
			l.Remove(l.Front())
			if l.Front() == nil {
				l.PushFront(i)
			} else {
				if i-l.Front().Value.(int) > max {
					max = i - l.Front().Value.(int)
				}
			}
		}
		//fmt.Println(i,max)
	}
	return max
}

//((()))
func LongestValidParentheses2(s string) int {
	//动态规划
	var max int
	var ds = make([]int, len(s))
	for i, item := range s {
		if i == 0 {
			continue
		}
		if item == '(' {
			//ds[i] = 0 可以省略
			continue
		}
		//上面的做了剪枝操作 避免代码分支太多可读性差
		//下面的情况都是）的情况
		if s[i-1] == '(' {
			//刚好配对
			//需要看i-2之前的能不能连起来
			if i-2 >= 0 {
				ds[i] = ds[i-2] + 2
			} else {
				ds[i] = 2
			}
			//是）的话去掉前面已经连起来的 再往前看一个是不是（能配对
		} else {
			//先把i=1的情况去掉 i=1时 ））直接不用处理
			if i == 1 {
				continue
			}

			if i-ds[i-1]-1 >= 0 && s[i-ds[i-1]-1] == '(' {

				ds[i] = ds[i-1] + 2
				if i-ds[i-1]-2 >= 0 {
					ds[i] = ds[i] + ds[i-ds[i-1]-2]
				}
			}

		}
		if ds[i] > max {
			max = ds[i]
		}
	}
	return max
}
